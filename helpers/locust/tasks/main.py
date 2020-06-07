import uuid

from datetime import datetime
from locust import HttpLocust, TaskSet, task


class MetricsTaskSet(TaskSet):
    _deviceid = None

    def on_start(self):
        self._deviceid = str(uuid.uuid4())

    @task(1)
    def get_request(self):
        self.client.get(
            '/get?foo1=bar1&foo2=bar2', timeout=30)

    @task(99)
    def post_metrics(self):
        self.client.post(
            "/post", {"deviceid": self._deviceid, "timestamp": datetime.now()})


class MetricsLocust(HttpLocust):
    task_set = MetricsTaskSet