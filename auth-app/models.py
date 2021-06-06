from app import db
from passlib.hash import pbkdf2_sha256 as sha256


class UserModel(db.Model):
    """
    User Model Class
    """

    __tablename__ = 'users'
    id = db.Column(db.Integer, primary_key=True)
    phone = db.Column(db.String(120), unique=True, nullable=False)
    name = db.Column(db.String(120), nullable=False)
    password = db.Column(db.String(120), nullable=False)
    role = db.Column(db.Integer, nullable=False)

    """
    Save user details in Database
    """
    def save_to_db(self):
        db.session.add(self)
        db.session.commit()

    """
    Find user by phone
    """
    @classmethod
    def find_by_phone(cls, phone):
        return cls.query.filter_by(phone=phone).first()

    """
    return all the user data in json form available in DB
    """
    @classmethod
    def return_all(cls):
        def to_json(x):
            return {
                'phone': x.phone,
                'password': x.password

            }

        return {'users': [to_json(user) for user in UserModel.query.all()]}

    """
    generate hash from password by encryption using sha256
    """
    @staticmethod
    def generate_hash(password):
        return sha256.hash(password)

    """
    Verify hash and password
    """
    @staticmethod
    def verify_hash(password, hash_):
        print(password)
        print(hash_)
        return sha256.verify(password, hash_)