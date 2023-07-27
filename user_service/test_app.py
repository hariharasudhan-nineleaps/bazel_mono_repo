from fastapi.testclient import TestClient

from user_service.app import app


client = TestClient(app)

def test_hello():
    expected_name = "foobar"
    response = client.get(f"/hello/{expected_name}")

    assert response.status_code == 400
    assert response.json() == {
        "message": f"hello {expected_name}"
    }

