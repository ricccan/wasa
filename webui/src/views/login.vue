<template>
    <div class="container mt-5">
        <div class="row justify-content-center">
            <div class="col-md-6">
                <div class="card">
                    <div class="card-header text-center">
                        <h3>Login</h3>
                    </div>
                    <div class="card-body">
                        <form @submit.prevent="doLogin">
                            <div class="mb-3">
                                <label for="nome" class="form-label text-center w-100">
                                    Nome utente
                                </label>
                                <input type="text" class="form-control" id="nome" v-model="nome" 
                                    placeholder="Inserisci il tuo nome utente" required >
                            </div>
                            <div class="d-grid">
                                <button type="submit" class="btn btn-primary">Login</button>
                            </div>
                        </form>
                        <p v-if="errorMessage" class="text-danger mt-3 text-center">
                            {{ errorMessage }}
                        </p>
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
            nome: "", // valore ottenuto dall'input
            errorMessage: "",
            id: null,
        };
    },
    methods: {
        async doLogin() { // doLogin (fa loggare l'utente)
            this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/session", {
                    l_name: this.nome,
                }); // crea un json che gli passa un nome
				this.some_data = response.data;
                this.id = response.data.Id;
                this.$router.push({ 
                    path: "/link1/" + this.id,
                    query: { username: this.nome,id: this.id },
                });

                localStorage.setItem("token", response.data.Token);
                localStorage.setItem("username", this.nome);
                localStorage.setItem("id", response.data.Id);
                
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
        },
    },
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