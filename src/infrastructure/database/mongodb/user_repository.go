package mongodb

import (
	"context"
	"os"

	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Client
}

// NewUserRepository cria uma nova instância do repositório MongoDB
func NewUserRepository(database *mongo.Client) user.UserRepositoryInterface {
	return &userRepository{
		db: database,
	}
}

func (ur *userRepository) Create(userDomain *user.UserDomain) (*user.UserDomain, *rest_err.RestErr) {
	collection := ur.db.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	logger.Info("Tentando inserir usuário no MongoDB")
	result, err := collection.InsertOne(context.Background(), userDomain)
	if err != nil {
		logger.Error("Erro ao inserir usuário no MongoDB", err)
		return nil, rest_err.NewInternalServerError("Erro ao criar usuário")
	}

	userDomain.ID = result.InsertedID.(primitive.ObjectID).Hex()
	logger.Info("Usuário inserido com sucesso no MongoDB")
	return userDomain, nil
}

func (ur *userRepository) FindByEmail(email string) (*user.UserDomain, *rest_err.RestErr) {
	collection := ur.db.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	userDomain := &user.UserDomain{}
	filter := bson.M{"email": email}

	err := collection.FindOne(context.Background(), filter).Decode(userDomain)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		logger.Error("Erro ao buscar usuário por email no MongoDB", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário")
	}

	return userDomain, nil
}

func (ur *userRepository) FindByID(id string) (*user.UserDomain, *rest_err.RestErr) {
	collection := ur.db.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("ID inválido")
	}

	userDomain := &user.UserDomain{}
	filter := bson.M{"_id": objectID}

	err = collection.FindOne(context.Background(), filter).Decode(userDomain)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		logger.Error("Erro ao buscar usuário por ID no MongoDB", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário")
	}

	return userDomain, nil
}

func (ur *userRepository) Update(id string, userDomain *user.UserDomain) *rest_err.RestErr {
	collection := ur.db.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rest_err.NewBadRequestError("ID inválido")
	}

	update := bson.M{
		"$set": bson.M{
			"name":       userDomain.Name,
			"email":      userDomain.Email,
			"age":        userDomain.Age,
			"updated_at": userDomain.UpdatedAt,
		},
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		logger.Error("Erro ao atualizar usuário no MongoDB", err)
		return rest_err.NewInternalServerError("Erro ao atualizar usuário")
	}

	return nil
}

func (ur *userRepository) Delete(id string) *rest_err.RestErr {
	collection := ur.db.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rest_err.NewBadRequestError("ID inválido")
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		logger.Error("Erro ao deletar usuário no MongoDB", err)
		return rest_err.NewInternalServerError("Erro ao deletar usuário")
	}

	return nil
}

func (ur *userRepository) FindAll() ([]*user.UserDomain, *rest_err.RestErr) {
	logger.Info("Iniciando busca de todos os usuários no MongoDB")
	collection := ur.db.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		logger.Error("Erro ao buscar usuários no MongoDB", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuários")
	}
	defer cursor.Close(context.Background())

	var users []*user.UserDomain
	if err := cursor.All(context.Background(), &users); err != nil {
		logger.Error("Erro ao decodificar usuários do MongoDB", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuários")
	}

	return users, nil
}
