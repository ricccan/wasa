<template>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css">
    <div class="container mt-5">
        <div class="row justify-content-center">
            <div class="col-md-6">
                <div class="card text-center">
                    <div class="card-header">
                        <h3>Profilo</h3>
                    </div>
                    <div class="card-body">
                        <img :src="photo" alt="User Photo" class="img-thumbnail mb-3 me-2"
                            style="width: 150px; height: 150px; object-fit: cover;" />
                        <button class="btn btn-primary" @click="editInfo"><i class="fas fa-pencil-alt" style="color: black;"></i></button>
                        <div class="d-flex align-items-center justify-content-center mt-3">
                            <h5 class="card-title me-2">{{ username }}</h5>
                            <button class="btn btn-primary" @click="cambiaNome"><i class="fas fa-pencil-alt" style="color: black;"></i></button>
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
            photo: null,
            id: localStorage.getItem("id"),
        };
    },
    methods: {
        editInfo() {
            alert("Edit Info functionality not implemented yet.");
        },
        cambiaNome() {
            alert("Cambia Nome functionality not implemented yet.");
        },
        async getPhoto() {
            this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/"+this.id+"/photo", {
                    headers:{ 
                        Authorization: "Bearer " + localStorage.getItem("token"),}

                }); // crea un json che gli passa un nome
                this.photo="data:image/jpeg;base64,"+response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
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
</style>