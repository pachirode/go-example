import socket
from gevent import monkey

print(f'Python socket: {socket.socket}')

monkey.patch_socket()
print(f'Gevent socket: {socket.socket}')

import multiprocessing

print(multiprocessing.cpu_count())

def cpu_count():
    return 100

multiprocessing.cpu_count = cpu_count
print(multiprocessing.cpu_count())
