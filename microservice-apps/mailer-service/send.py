import pika
import os

rmq_host = os.getenv('RABBIT_HOST', 'localhost')

connection = pika.BlockingConnection(pika.ConnectionParameters(rmq_host))
channel = connection.channel()

channel.queue_declare(queue='notification')

channel.basic_publish(exchange='', routing_key='notification',
                      body='{"name":"John", "email":"test@example.com"}')
print(" [x] Sent 'Example body'")
connection.close()
