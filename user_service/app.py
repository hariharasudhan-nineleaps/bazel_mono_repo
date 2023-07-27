import sys
import platform

# external packages
from fastapi import FastAPI


app = FastAPI()

@app.get("/hello/{name}")
async def hello(name: str):
    return {"message": f"hello {name}"}



