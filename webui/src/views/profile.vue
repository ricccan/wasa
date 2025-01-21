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
                            <button class="edit-button mb-2" @click="showPopup1 = true"><i class="fas fa-pencil-alt"
                                    style="color: black;"></i></button>
                            <!-- popup nel caso prema su cambia nome -->
                            <div v-if="showPopup1" class="popup-overlay" @click.self="closePopup">
                                <div class="popup-content">
                                    <h2>New username</h2>
                                    <input type="text" v-model="newUsername" placeholder="Type here..." />
                                    <div class="popup-actions">
                                        <button @click="cambiaNome">Submit</button>
                                        <button @click="closePopup">Close</button>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="position-relative">
                            <!-- sezione foto profilo -->
                            <img :src="photo" alt="User Photo" class="img-thumbnail" />
                            <button class="edit-button position-absolute bottom-0 end-0 m-2 " @click="showPopup = true">
                                <i class="fas fa-pencil-alt"></i>
                            </button>
                        </div>
                        <!-- popup nel caso prema su cambia foto -->
                        <div v-if="showPopup" class="popup-overlay" @click.self="closePopup">
                            <div class="popup-content">
                                <h2>New photo</h2>
                                <input type="file" @change="onFileChange" accept="image/*" />

                                <!-- Image Preview -->
                                <div v-if="imagePreview" class="image-preview">
                                    <img :src="imagePreview" alt="Selected photo preview" />
                                </div>
                                <div class="popup-actions">
                                    <button @click="cambiaFoto">Submit</button>
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
        };
    },
    methods: {
        editInfo() {
            alert("Edit Info functionality not implemented yet.");
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
            } catch (e) {
                this.errormsg = e.toString();
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
    },
    mounted() {
        this.getPhoto(); // per chiamare le funzioni istantaneamente al caricamento dalla pagnia
    }
};
</script>

<style scoped>
.container {
    margin-top: 100px;
}

.card {
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.edit-icon-container {
  position: relative;
  display: inline-block;
  width: 40px;
  height: 40px;
  background-color: white;
  border: 2px solid #007bff;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease, transform 0.2s ease;
}

/* Pencil Icon */
.edit-icon {
  color: #007bff;
  font-size: 16px;
  transition: color 0.3s ease;
}

/* Hover Effects */
.edit-icon-container:hover {
  background-color: #007bff;
  transform: scale(1.1);
}

.edit-icon-container:hover .edit-icon {
  color: white;
}

</style>