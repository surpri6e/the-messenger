import socket
import threading


class ChatClient:
    def __init__(self, host='localhost', port=5555):
        self.client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.client.connect((host, port))
        self.nickname = None

    def receive(self):
        """Получение сообщений от сервера"""
        while True:
            try:
                message = self.client.recv(1024).decode('utf-8')

                if message == "NICK":
                    self.client.send(self.nickname.encode('utf-8'))
                else:
                    print(message)
            except:
                print("Произошла ошибка!")
                self.client.close()
                break

    def write(self):
        """Отправка сообщений на сервер"""
        while True:
            try:
                message = input('')
                if message.lower() == '/quit':
                    self.client.close()
                    break

                full_message = f"{self.nickname}: {message}"
                self.client.send(full_message.encode('utf-8'))
            except:
                break

    def start(self):
        """Запуск клиента"""
        self.nickname = input("Выберите никнейм: ")

        receive_thread = threading.Thread(target=self.receive)
        receive_thread.start()

        write_thread = threading.Thread(target=self.write)
        write_thread.start()


if __name__ == "__main__":
    client = ChatClient()
    client.start()