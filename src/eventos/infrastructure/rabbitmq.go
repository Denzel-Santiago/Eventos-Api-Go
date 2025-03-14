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

// PublishTicketPurchaseMessage envía un mensaje con la compra de boletos
func PublishTicketPurchaseMessage(ch *amqp.Channel, queueName string, eventID string, ticketsSold int) error {
	message := TicketPurchaseMessage{
		EventID:     eventID,
		TicketsSold: ticketsSold,
	}

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error convirtiendo mensaje a JSON: %v", err)
	}

	err = ch.Publish(
		"",        // Exchange ("" significa que va directo a la cola)
		queueName, // Routing key (nombre de la cola)
		false,     // Mandatory
		false,     // Immediate
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
