import socket
import threading


class ChatServer:
    def __init__(self, host='localhost', port=5555):
        self.server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.server.bind((host, port))
        self.server.listen()

        self.clients = []
        self.nicknames = []

        print(f"Сервер запущен на {host}:{port}")

    def broadcast(self, message, sender_client=None):
        """Отправка сообщения всем клиентам"""
        for client in self.clients:
            if client != sender_client:
                try:
                    client.send(message)
                except:
                    self.remove_client(client)

    def remove_client(self, client):
        """Удаление отключившегося клиента"""
        if client in self.clients:
            index = self.clients.index(client)
            self.clients.remove(client)
            client.close()
            nickname = self.nicknames[index]
            self.nicknames.remove(nickname)
            self.broadcast(f"{nickname} покинул чат!".encode('utf-8'))

    def handle_client(self, client):
        """Обработка сообщений от конкретного клиента"""
        while True:
            try:
                message = client.recv(1024)
                if not message:
                    break
                self.broadcast(message, client)
            except:
                self.remove_client(client)
                break

    def receive(self):
        """Прием новых подключений"""
        while True:
            client, address = self.server.accept()
            print(f"Подключен {str(address)}")

            # Запрашиваем никнейм
            client.send("NICK".encode('utf-8'))
            nickname = client.recv(1024).decode('utf-8')

            self.nicknames.append(nickname)
            self.clients.append(client)

            print(f"Никнейм: {nickname}")
            self.broadcast(f"{nickname} присоединился к чату!".encode('utf-8'))
            client.send("Вы подключены к серверу!".encode('utf-8'))

            # Запускаем поток для обработки клиента
            thread = threading.Thread(target=self.handle_client, args=(client,))
            thread.start()


if __name__ == "__main__":
    server = ChatServer()
    server.receive()