<template>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css">
    <div class="container mt-5">
        <div class="row justify-content-center">
            <div class="col-md-6">
                <div class="card text-center">
                    <div class="card-header">
                        <h3>Profile</h3>
                    </div>
                    <!-- sezione centrale -->
                    <div class="card-body">
                        <div class="d-flex align-items-center justify-content-center mt-2">
                            <!-- sezione nome utente -->
                            <h5 class="card-title me-2 mb-2 ">{{ username }}</h5>
                            <button class="btn btn-outline-light mb-2" @click="showPopup1 = true">
                                <i class="fas fa-pencil-alt" style="color: black;"></i>
                            </button>
                            <!-- popup nel caso prema su cambia nome -->
                            <div v-if="showPopup1" class="popup-overlay" @click.self="closePopup">
                                <div class="popup-content">
                                    <h2>New username</h2>
                                    <input type="text" v-model="newUsername" placeholder="Type here..." />
                                    <div class="popup-actions">
                                        <button @click="cambiaNome" v-if="newUsername">Submit</button>
                                        <button @click="closePopup">Close</button>
                                    </div>
                                    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
                                </div>
                            </div>
                        </div>
                        <div class="position-relative">
                            <!-- sezione foto profilo -->
                            <img :src="photo" alt="User Photo" class="img-thumbnail" id="foto1" />
                            <button class="edit-button position-absolute bottom-0 end-0 m-2 " @click="showPopup = true">
                                <i class="fas fa-pencil-alt"></i>
                            </button>
                        </div>
                        <button class="btn btn-primary mt-3" @click="home">Home</button>
                        <!-- popup nel caso prema su cambia foto -->
                        <div v-if="showPopup" class="popup-overlay" @click.self="closePopup">
                            <div class="popup-content">
                                <h2>New photo</h2>
                                <div class="file-upload-container">
                                    <!-- File Input -->
                                    <input type="file" @change="onFileChange" accept="image/*" />

                                    <!-- Remove Photo Button -->
                                    <button v-if="imagePreview" @click="removePhoto()" class="remove-button">
                                        <i class="fas fa-trash-alt"></i>
                                    </button>
                                </div>

                                <!-- Image Preview -->
                                <div v-if="imagePreview" class="image-preview">
                                    <img :src="imagePreview" alt="Selected photo preview" />
                                </div>
                                <div class="popup-actions">
                                    <button @click="cambiaFoto" v-if="imagePreview" >Submit</button>
                                    <button @click="closePopup">Close</button>
                                </div>
                            </div>
                        </div>

                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            username: localStorage.getItem("username"), // prendo la variabile username
            photo: null, // variabile foto
            id: localStorage.getItem("id"), // prendo l'id dell'utente che ha fatto il login
            newUsername: "",
            showPopup1: false, // variabili per far aprire i popup
            showPopup: false,
            textInput: "",
            selectedFile: null,
            imagePreview: null,
            errormsg: null,
        };
    },
    methods: {
        removePhoto() {
            this.selectedFile = null;
            this.imagePreview = null;
        },
        async cambiaFoto() {
            this.loading = true;
            this.errormsg = null;
            try {
                let formData = new FormData();
                formData.append("photo", this.selectedFile);

                let response = await this.$axios.post("/users/" + this.id + "/photo", formData, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token"),
                        "Content-Type": "multipart/form-data",
                    }
                });
                this.getPhoto();
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async cambiaNome() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post("/users/" + this.id + "/username", {
                    un_name: this.newUsername,
                },
                    {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("token"),
                        }



                    }); // crea un json che gli passa un nome
                this.username = response.data.un_username;
                localStorage.setItem("username", this.username);
                this.showPopup1 = false;
            } catch (e) {
                this.errormsg = e.response.data;
            }
            this.loading = false;
            
        },
        async getPhoto() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.get("/users/" + this.id + "/photo", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token"),
                    }

                }); // crea un json che gli passa un nome
                this.photo = "data:image/jpeg;base64," + response.data;
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        closePopup() {
            this.showPopup = false;
            this.showPopup1 = false;
        },
        onFileChange(event) {
            const file = event.target.files[0];
            if (file) {
                this.selectedFile = file;

                // Create a preview of the image
                const reader = new FileReader();
                reader.onload = (e) => {
                    this.imagePreview = e.target.result;
                };
                reader.readAsDataURL(file);
            }
        },
        home() {
            this.$router.push({ path: "/link1" });
        }
    },


    mounted() {
        this.getPhoto(); // per chiamare le funzioni istantaneamente al caricamento dalla pagnia
    }
};
</script>

<style scoped>


.card {
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.file-upload-container {
  display: flex;
  align-items: center; /* Vertically center items */
  gap: 10px; /* Add spacing between the input and button */
}

.remove-button {
  padding: 10px 15px;
  border: none;
  background-color: #e74c3c;
  color: white;
  border-radius: 5px;
  cursor: pointer;
}

.remove-button:hover {
  background-color: #c0392b;
}



</style>