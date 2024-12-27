from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()


class User(db.Model):
    __tablename__ = 'users'

    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(50), nullable=False)
    lastname = db.Column(db.String(50), nullable=False)
    email = db.Column(db.String(50), nullable=False)
    created_at = db.Column(db.DateTime(), nullable=False,
                           default=db.func.current_timestamp())

    @classmethod
    def create(cls, username: str, lastname: str, email: str):
        """Crea el usuario por medio de un metodo estatico

        Args:
            username (string): Nombre del usuario
            lastname (string): Apellido del usuario
            email (string): Email

        Returns:
            User: La clase generada desde base de datos
        """
        user = User(username=username, lastname=lastname, email=email)
        return user.save()

    def update(self):
        """Se actualiza el registro en base de datos, se retorna el resultado"""
        self.save()
        return self

    def update_field(self, field: str, value):
        """Actualiza la clase temporalmente

        Args:
            field (str): nombre del campo
            value (str): valor

        Returns:
            _type_: _description_
        """
        if field == "username":
            self.username = value if value is not None else self.username
        if field == "lastname":
            self.lastname = value if value is not None else self.lastname
        if field == "email":
            self.email = value if value is not None else self.email

        return self

    def save(self):
        """Crea el usuario y lo guarda en base de datos

        Returns:
            User: usuario de base de datos
        """
        try:
            db.session.add(self)
            db.session.commit()
            return self
        except:
            return False

    def delete(self):
        """Borramos el registro"""
        try:
            db.session.delete(self)
            db.session.commit()
            return True
        except:
            return False

    def to_json(self):
        """Convierte el objeto User a un diccionario JSON

        Returns:
            dict: Diccionario con los datos del usuario
        """
        return {
            'id': self.id,
            'username': self.username,
            'lastname': self.lastname,
            'email': self.email,
            'created_at': self.created_at.isoformat()
        }
