package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseConnection *mongo.Client
}

func NewUserRepository(database *mongo.Client) user.UserRepositoryInterface {
	return &userRepository{
		databaseConnection: database,
	}
}

func (ur *userRepository) Create(userDomain *user.UserDomain) (*user.UserDomain, *rest_err.RestErr) {
	collection := ur.databaseConnection.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	userDomain.CreatedAt = time.Now()
	userDomain.UpdatedAt = time.Now()

	logger.Info("Attempting to insert user into MongoDB")
	result, err := collection.InsertOne(context.Background(), userDomain)
	if err != nil {
		logger.Error("Error inserting user into MongoDB", err)
		return nil, rest_err.NewInternalServerError("Error trying to create user")
	}

	userDomain.ID = result.InsertedID.(primitive.ObjectID).Hex()
	logger.Info("User successfully inserted into MongoDB")
	return userDomain, nil
}

func (ur *userRepository) FindByEmail(email string) (*user.UserDomain, *rest_err.RestErr) {
	logger.Info("Iniciando busca de usuário por email no MongoDB")
	collection := ur.databaseConnection.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	userDomain := &user.UserDomain{}
	filter := bson.M{"email": email}

	err := collection.FindOne(context.Background(), filter).Decode(userDomain)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("Usuário não encontrado no MongoDB", err)
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		logger.Error("Erro ao buscar usuário por email no MongoDB", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário")
	}

	logger.Info("Usuário encontrado com sucesso no MongoDB")
	return userDomain, nil
}

func (ur *userRepository) FindByID(id string) (*user.UserDomain, *rest_err.RestErr) {
	logger.Info("Iniciando busca de usuário por ID no MongoDB")
	collection := ur.databaseConnection.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Erro ao converter ID", err)
		return nil, rest_err.NewBadRequestError("ID inválido")
	}

	userDomain := &user.UserDomain{}
	filter := bson.M{"_id": objectID}

	err = collection.FindOne(context.Background(), filter).Decode(userDomain)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("Usuário não encontrado no MongoDB", err)
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		logger.Error("Erro ao buscar usuário por ID no MongoDB", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário")
	}

	logger.Info("Usuário encontrado com sucesso no MongoDB")
	return userDomain, nil
}

func (ur *userRepository) Update(id string, userDomain *user.UserDomain) *rest_err.RestErr {
	logger.Info("Iniciando atualização de usuário no MongoDB")
	collection := ur.databaseConnection.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Erro ao converter ID", err)
		return rest_err.NewBadRequestError("ID inválido")
	}

	// Cria o objeto de atualização apenas com os campos que foram fornecidos
	updateFields := bson.M{}

	// Só inclui os campos que foram fornecidos no userDomain
	if userDomain.Name != "" {
		updateFields["name"] = userDomain.Name
	}
	if userDomain.Email != "" {
		updateFields["email"] = userDomain.Email
	}
	if userDomain.Age > 0 {
		updateFields["age"] = userDomain.Age
	}
	if userDomain.Password != "" {
		updateFields["password"] = userDomain.Password
	}

	// Só atualiza se houver campos para atualizar
	if len(updateFields) > 0 {
		userDomain.UpdatedAt = time.Now()
		updateFields["updated_at"] = userDomain.UpdatedAt

		update := bson.M{
			"$set": updateFields,
		}

		_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
		if err != nil {
			logger.Error("Erro ao atualizar usuário no MongoDB", err)
			return rest_err.NewInternalServerError("Erro ao atualizar usuário")
		}

		logger.Info("Usuário atualizado com sucesso no MongoDB")
	} else {
		logger.Info("Nenhum campo para atualizar")
	}

	return nil
}

func (ur *userRepository) Delete(id string) *rest_err.RestErr {
	logger.Info("Iniciando exclusão de usuário no MongoDB")
	collection := ur.databaseConnection.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Erro ao converter ID", err)
		return rest_err.NewBadRequestError("ID inválido")
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		logger.Error("Erro ao deletar usuário no MongoDB", err)
		return rest_err.NewInternalServerError("Erro ao deletar usuário")
	}

	logger.Info("Usuário deletado com sucesso no MongoDB")
	return nil
}
