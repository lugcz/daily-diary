# -*- coding: utf-8 -*-
from fastapi import FastAPI

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.get("/new_route")
async def new_route():
    return {"key": "value"}