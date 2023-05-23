package userController

import (
	dbconfig "go-gin-api/config/dbConfig"
	"go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reads(ctx *gin.Context){
  var user []models.User

  dbconfig.DB.Find(&user)

  ctx.JSON(200, gin.H{
    "message":"Data user tersedia",
    "data": user,
  })
}

func Read(ctx *gin.Context) {
  var user models.User

  userId := ctx.Param("id")

  if userId == "" {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"Id tidak boleh kosong",
    })
    return 
  }

  if err := dbconfig.DB.Where("id = ?", userId).First(&user).Error; err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"Data  user tidak ditemukan",
    })
    return
  }

  ctx.JSON(http.StatusOK, gin.H{
    "message":"Data user tersedia",
    "data": user,
  })
}

func Create(ctx *gin.Context){
  userReq := new(models.UserReq)

  if err := ctx.BindJSON(&userReq); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"Gagal parsing JSON",
    })
    return
  }
  
  var user models.User
  user.Name = userReq.Name
  user.Address = userReq.Address
  

  if err := dbconfig.DB.Create(&user).Error; err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "message":"Internal server Error",
    })
    return
  }

  ctx.JSON(http.StatusOK, gin.H{
    "message":"Data berhasil dibuat",
    "data": user,
  })
}

func Update(ctx *gin.Context){
  userReq := new(models.UserReq)

  if err := ctx.ShouldBindJSON(&userReq); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"Gagal parsing JSON",
    })
    return
  }

  userId := ctx.Param("id")

  if userId == "" {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"id tidak boleh kosong!",
    })
    return
  }

  updateUser := models.User{}

  if err := dbconfig.DB.First(&updateUser, "id = ? ", userId).Error; err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"Data user tidak ditemukan",
    })
    return
  }

  updateUser.Name = userReq.Name
  updateUser.Address = userReq.Address

  if err := dbconfig.DB.Model(&updateUser).Error; err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message": "error models update user",
    })
    return
  }

  result := dbconfig.DB.Where("id = ?", userId).Updates(&updateUser)
  if result.Error != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"Id user tidak ditemukan!",
    })
    return
  }

  if result.RowsAffected == 0 {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"Data tida ditemukan",
    })
    return
  }

  ctx.JSON(http.StatusOK, gin.H{
    "message":"Data berhasil diupdate",
    "data": updateUser,
  })
}

func Delete(ctx *gin.Context){
  user := models.User{}
  userId := ctx.Param("id")

  if err := dbconfig.DB.First(&user, userId).Delete(&user).Error; err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message":"Gagal hapus data user",
    })
    return
  }

  ctx.JSON(http.StatusOK, gin.H{
    "message":"Data user berhasil dihapus",
  })
}

