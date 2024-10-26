package rabittmq

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/jhonM8a/worker-evaluacion/config"
	"github.com/streadway/amqp"
)

var (
	once     sync.Once
	instance *RabbitMQ
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	mu      sync.Mutex
}

type Message struct {
	IdEvaluation int  `json:"idEvaluation"`
	IsValid      bool `json:"isValid"`
}

func getInstance() (*RabbitMQ, error) {
	var err error
	once.Do(func() {
		conf := config.LoadConfRabbitMq()
		// Corrección de la URL de conexión
		conn, connErr := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Pass, conf.Host, conf.Port))
		if connErr != nil {
			err = connErr
			fmt.Printf("Error al conectarme: %v\n", err)

			return
		}

		ch, chErr := conn.Channel()
		if chErr != nil {
			err = chErr
			fmt.Printf("Error al crear canal: %v\n", err)

			return
		}

		instance = &RabbitMQ{
			conn:    conn,
			channel: ch,
		}
	})

	if err != nil {
		return nil, err
	}
	return instance, nil
}

// Función para encolar mensajes (estructura Message)
func Enqueue(msg Message) error {

	conf := config.LoadConfRabbitMq()

	// Obtener instancia de RabbitMQ (singleton)
	rmq, err := getInstance()
	if err != nil {
		fmt.Printf("Error --->: %v\n", err)

		return err
	}

	rmq.mu.Lock()
	defer rmq.mu.Unlock()

	// Serializar el mensaje a JSON
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Mensaje: %v\n", body)

	// Declarar la cola si no existe
	q, err := rmq.channel.QueueDeclare(
		conf.NameQueueEva, // nombre de la cola
		false,             // durable
		false,             // auto-delete
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		fmt.Printf("Error al encolar: %v\n", err)

		return err
	}

	// Publicar mensaje
	err = rmq.channel.Publish(
		"",     // exchange
		q.Name, // clave de enrutamiento
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		fmt.Printf("Error al encolar2: %v\n", err)

		return err
	}

	log.Printf("Mensaje encolado: %s", body)
	return nil
}

// Función para cerrar la conexión de RabbitMQ (se usará manualmente cuando sea necesario)
func (r *RabbitMQ) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
