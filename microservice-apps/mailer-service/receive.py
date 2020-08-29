import pika
import os
import sys
import smtplib
from socket import gaierror
import json


def main():
    rmq_host = os.getenv('RABBIT_HOST', 'localhost')
    mail_username = os.getenv('MAIL_USERNAME', 'username')
    mail_password = os.getenv('MAIL_PASSWORD', 'password')
    smtp_host = os.getenv('SMTP_HOST', 'localhost')
    smtp_port = os.getenv('SMTP_PORT', '25')

    connection = pika.BlockingConnection(pika.ConnectionParameters(rmq_host))
    channel = connection.channel()

    channel.queue_declare(queue='notification')

    def callback(ch, method, properties, body):
        print(" [x] Received %r" % body)

        msg = json.loads(body)
        receiver_name = msg["name"]
        receiver_email = msg["email"]

        sender = "no-reply <no-reply@microservice.id>"
        receiver = f"{receiver_name} <{receiver_email}>"

        # type your message: use two newlines (\n) to separate the subject from the message body, and use 'f' to  automatically insert variables in the text
        message = f"""\
        Subject: New Login Detected
        To: {receiver}
        From: {sender}

        Hi {receiver_name}! We detected new login for your account."""

        try:
            # send your message with credentials specified above
            with smtplib.SMTP(smtp_host, int(smtp_port)) as server:
                server.login(mail_username, mail_password)
                server.sendmail(sender, receiver, message)
            # tell the script to report if your message was sent or which errors need to be fixed
            print(f' [v] Success sent to {receiver_email}')
        except (gaierror, ConnectionRefusedError):
            print('Failed to connect to the server. Bad connection settings?')
        except smtplib.SMTPServerDisconnected:
            print('Failed to connect to the server. Wrong user/password?')
        except smtplib.SMTPException as e:
            print('SMTP error occurred: ' + str(e))

    channel.basic_consume(
        queue='notification', on_message_callback=callback, auto_ack=True)

    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
