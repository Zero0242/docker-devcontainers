from os import getenv
from typing import Optional
from dotenv import load_dotenv
from flask import Flask, jsonify, request
from src.config import config
from src.models import db, User


def create_app(enviroment):
    """Configuracion del server, crea la app inicial"""
    flask_app = Flask(__name__)
    flask_app.config.from_object(enviroment)
    with flask_app.app_context():
        db.init_app(flask_app)
        db.create_all()
    return flask_app


app = create_app(config['development'])


@app.get("/")
def get_home():
    return "Hola mundo desde Flask"


@app.get("/api/v1/users/<uid>")
def get_user(uid):
    user: Optional[User] = User.query.filter_by(id=uid).first()
    if user is None:
        return jsonify({"ok": False, "message": f"User with id {uid} not found"}), 404
    return jsonify(user.to_json())


@app.get("/api/v1/users")
def get_users():
    users = [user.to_json() for user in User.query.all()]
    return jsonify(users), 200


@app.post("/api/v1/users")
def create_user():
    json = request.get_json(force=True)
    for field in ["username", "lastname", "email"]:
        if json.get(field) is None:
            return jsonify({"ok": False, "message": f"Field {field} is missing!"}), 400

    user = User.create(json["username"], json["lastname"], json["email"])
    return jsonify({"ok": True, "user": user.to_json()}), 201


@app.patch("/api/v1/users/<uid>")
def update_user(uid):
    user: Optional[User] = User.query.filter_by(id=uid).first()
    if user is None:
        return jsonify({"ok": False, "message": f"User with id {uid} not found"}), 404

    json = request.get_json(force=True)
    for field in ["username", "lastname", "email"]:
        user = user.update_field(field, json[field])

    user.update()

    return jsonify(user.to_json())


@app.delete("/api/v1/users/<uid>")
def delete_user(uid):
    user: Optional[User] = User.query.filter_by(id=uid).first()
    if user is None:
        return jsonify({"ok": False, "message": f"User with id {uid} not found"}), 404
    user.delete()
    return jsonify({"ok": True, "message": f"Usuario {uid} eliminado"})


if __name__ == "__main__":
    load_dotenv()
    # FLASK_DEBUG es el valor de debug
    app.run(debug=True, host="0.0.0.0", port=getenv("PORT", default="3000"))
