package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// TicketPurchaseMessage representa el mensaje que se enviará a la cola
type TicketPurchaseMessage struct {
	EventID     string `json:"event_id"`
	TicketsSold int    `json:"tickets_sold"`
}

// ConnectRabbitMQ establece la conexión con RabbitMQ
func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://Denzel:Desz117s@18.211.110.229:5672/")
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %s", err)
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error creando canal: %s", err)
		return nil, nil, err
	}

	fmt.Println("✅ Conectado a RabbitMQ correctamente")
	return conn, ch, nil
}

// DeclareQueue crea una cola en RabbitMQ
func DeclareQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)

	if err != nil {
		log.Fatalf("Error declarando la cola: %s", err)
		return q, err
	}

	fmt.Printf("✅ Cola '%s' verificada correctamente\n", queueName)
	return q, nil
}


func PublishTicketPurchaseMessage(ch *amqp.Channel, message map[string]interface{}) error {
    queueName := "queue"

    body, err := json.Marshal(message)  // Convertir el JSON dinámico a string
    if err != nil {
        return fmt.Errorf("error convirtiendo mensaje a JSON: %v", err)
    }

    err = ch.Publish(
        "",        // Intercambio vacío
        queueName, // Nombre de la cola
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )

    if err != nil {
        return fmt.Errorf("error enviando mensaje a RabbitMQ: %v", err)
    }

    fmt.Printf("📩 Mensaje enviado a la cola '%s': %s\n", queueName, body)
    return nil
}
