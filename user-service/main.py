import sys
import platform

# external packages
import uvicorn
from fastapi import FastAPI


app = FastAPI()

@app.get("/")
def index():
    bazel_python_path = f'Python executable used by Bazel is: {sys.executable} <br/><br/>'
    bazel_python_version = f'Python version used by Bazel is: {platform.python_version()} <br/><br/>'

    output_str = (
        bazel_python_path
        + bazel_python_version
    )

    return {"output": output_str}


@app.get("/api/info")
async def get_info():
    return {"info": "active"}


if __name__ == '__main__':
    uvicorn.run("user-service.main:app", port=5000)


