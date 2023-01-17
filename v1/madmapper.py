import argparse
import time

from pythonosc import udp_client

last_theme = None

def send_theme(data):
    parser = argparse.ArgumentParser()
    parser.add_argument("--ip", default="192.168.1.205",
        help="The ip of the OSC server")
    parser.add_argument("--port", type=int, default=8011,
        help="The port the OSC server is listening on")
    args = parser.parse_args()

    client = udp_client.SimpleUDPClient(args.ip, args.port)

    address = "/" + data

    client.send_message(address, 0.5)


    # client.close()
