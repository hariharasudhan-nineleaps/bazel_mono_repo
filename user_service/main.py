import uvicorn

if __name__ == '__main__':
    uvicorn.run("user_service.app:app", port=5000)