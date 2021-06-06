from flask_restful import Resource, reqparse
from models import UserModel
import string, random
from flask_jwt_extended import (
    create_access_token,
    create_refresh_token,
    jwt_required,
    get_jwt_identity,
)

parser = reqparse.RequestParser()
parser.add_argument('name')
parser.add_argument('phone')
parser.add_argument('role', help='role cannot be blank', required=False)
parser.add_argument('password')

class UserRegistration(Resource):
    """
    User Registration Api
    """

    def post(self):
        data = parser.parse_args()
        phone = data['phone']
        name = data['name']
        role = data['role']

        random = password_generator()
        print(parser)
        # Checking that user is already exist or not
        if UserModel.find_by_phone(phone):

            return {
                'code': 500,
                'message': f'User {name} with number phone {phone} is already exists'
            }

        # create new user
        new_user = UserModel(
            phone=phone,
            name=name,
            role=role,
            password=UserModel.generate_hash(random)
        )
        print(new_user)
        try:
            
            # Saving user in DB and Generating Access and Refresh token
            new_user.save_to_db()
            access_token = create_access_token(identity=phone)
            refresh_token = create_refresh_token(identity=phone)
            
            return {
                'code': '00',
                'message': f'User {name} was created',
                'data': {
                    'access_token': access_token,
                    'refresh_token': refresh_token,
                    'password': random,
                }
            }
        
        except:
        
            return {
                'code': 500,
                'message': 'Something went wrong'
            }, 500


class UserLogin(Resource):
    """
    User Login Api
    """

    def post(self):
        data = parser.parse_args()
        phone = data['phone']
        password = data['password']

        # Searching user by phone
        current_user = UserModel.find_by_phone(phone)
        
        # user does not exists
        if not current_user:
        
            return {'message': f'User {phone} doesn\'t exist'}
        
        print("password sekarang",current_user.password)
        print("password dikirim",password)
        # user exists, comparing password and hash

        if UserModel.verify_hash(password, current_user.password):
            # generating access token and refresh token
            access_token = create_access_token(identity=phone)
            refresh_token = create_refresh_token(identity=phone)
        
            return {
                'message': f'Logged with phone number {phone}',
                'access_token': access_token,
                'refresh_token': refresh_token
            }
        else:
        
            return {
                'code': 500,
                'message': "Incorrect username or password"
            }

class Profile(Resource):
    
    @jwt_required()
    def get(self):
        """
        return all user api
        """
        phone = get_jwt_identity()

        current_user = UserModel.find_by_phone(phone)
        print(current_user)
        return {
            'code': '00',
            'message': "Success",
            'data': {
                'phone': current_user.phone,
                'name': current_user.name,
                'role': current_user.role
            }
        }

def password_generator(size=4, chars=string.ascii_uppercase + string.digits):
    return ''.join(random.choice(chars) for _ in range(size))