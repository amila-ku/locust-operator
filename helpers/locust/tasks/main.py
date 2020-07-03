import uuid


import random
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(5, 9)

    @task
    def index_page(self):
        self.client.get("/get?foo1=bar1")
        self.client.get("/get?foo2=bar2")

    @task(3)
    def view_item(self):
        item_id = random.randint(1, 10000)
        self.client.get(f"/get?id={item_id}", timeout=30)

    def on_start(self):
        self.client.post("/post", {"deviceid": self._deviceid, "timestamp": datetime.now()})