import ssl
from paho import mqtt
import paho.mqtt.client as paho
import paho.mqtt.publish as publish
import os
from dotenv import load_dotenv
import json

load_dotenv()

username = os.getenv("HIVE_MQ_USER")
password = os.getenv("HIVE_MQ_PASS")
hostname = os.getenv("HIVE_MQ_URL")

sslSettings = ssl.SSLContext(mqtt.client.ssl.PROTOCOL_TLS)
auth = {'username': username, 'password': password}

data = {'temperature': 20, 'humidity': 0.9, "device_id":"fd6c:6361:7a9f:0:6c90:5591:8ea8:xxxa"}
topic = "measures"

payload = json.dumps(data)


publish.single(topic=topic, payload=payload, hostname=hostname, port=8883, auth=auth,
                 tls=sslSettings, protocol=paho.MQTTv31)
