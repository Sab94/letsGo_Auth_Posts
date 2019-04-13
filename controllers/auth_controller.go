/*
|--------------------------------------------------------------------------
| Authentication Controller
|--------------------------------------------------------------------------
|
| GetCredentials works on oauth2 Client Credentials Grant and returns CLIENT_ID, CLIENT_SECRET
| GetToken takes CLIENT_ID, CLIENT_SECRET, grant_type, scope and returns access_token and some other information
*/

package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"github.com/google/uuid"
	"github.com/posts-api/database"
	"github.com/posts-api/helpers"
	"github.com/posts-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
)
var clientStore = store.NewClientStore()
var manager = manage.NewDefaultManager()

func AuthInit() {
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	manager.MapClientStorage(clientStore)

	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)
}

func GetCredentials(c *gin.Context) {
	clientId := uuid.New().String()
	clientSecret := uuid.New().String()
	err := clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: "http://localhost:8000",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
	c.Done()
}

func GetToken(c *gin.Context) {
	ginserver.HandleTokenRequest(c)
}

func Verify (c *gin.Context) {
	ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
	if exists {
		c.JSON(200, ti)
		return
	}
	c.String(200, "not found")
}

func Register (c *gin.Context) {
	a := types.User{}
	ctx := context.Background()
	collection := database.DB.Collection("users")
	err := c.BindJSON(&a)
	a.Password,_ = Generate(a.Password)
	a.Id = primitive.NewObjectID()

	if err != nil {
		log.Println(err)
		c.Abort()
	}
	_, err = collection.InsertOne(ctx, a)
	if err != nil {
		log.Println(err)
		c.Abort()
	}
	c.JSON(200, a)
	c.Done()
}

func Login (c *gin.Context) {
	a := types.LoginRequest{}

	collection := database.DB.Collection("users")
	err := c.BindJSON(&a)

	user := types.User{}
	err = collection.FindOne(context.Background(), bson.M{"email": a.Email}).Decode(&user)

	if err != nil {
		log.Println(err)
		c.Abort()
	}
	loginError := Compare(user.Password, a.Password)

	if loginError != nil {
		log.Println(err)
		c.JSON(helpers.ErrorMessage(err, types.ErrLogin))
		c.Done()
	} else {
		clientId, clientSecret := getCredentialsForLogin(user.Id, user.Email, user.Password)

		res := types.LoginResponse{}
		res.CLIENT_ID = clientId
		res.CLIENT_SECRET = clientSecret
		res.User = user
		c.JSON(200, res)
		c.Done()
	}


}

//Generate a salted hash for the input string
func Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)

	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func getCredentialsForLogin (user_id primitive.ObjectID, client_id string, client_secret string) (string, string) {
	err := clientStore.Set(client_id, &models.Client{
		ID:     client_id,
		Secret: client_secret,
		UserID: user_id.String(),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return client_id, client_secret
}
