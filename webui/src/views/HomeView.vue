<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			token: localStorage.getItem("token"), // prende il valore del token dato dal login
			username: localStorage.getItem("username"), // prende il valore dell'username dato nel login
			chats: {}, // json che conterrà la lista delle chat
			id: localStorage.getItem("id"), // id utente collegato allo username
			dynamicData: {}, // json che utilizzo insieme a chats per gestire le conversazioni
			showPopup: false, // popup che apre la sezione "crea gruppo"
			showPopup1: false, // popup che apre la sezione "nuova chat"
			newGroupname: null, // valore che contiene il nome da dare al gruppo che sta venendo creato
			newUser: null, // valore dello username dell'utente con cui si vuole creare una nuova conversazione
			messages: {}, // json a cui passo i messaggi
			dynamicData2: {}, // json che utilizzo insieme a chats per gestire i messaggi
			nomi: {} // variabile a cui do il valore dei nomi nei messaggi
		}
	},
	methods: {
		async newPage() { // funzione che porta alla pagina del profilo
			this.loading = true;
			this.errormsg = null;
			try {
				this.$router.push({ path: "/profile" }); // reindirizza alla pagina del profilo
			} catch (e) {
				this.errormsg = e.toString(); // gestisco errori
			}
			this.loading = false;
		},
		async login() { // funzione che porta alla pagina del login
			this.loading = true;
			this.errormsg = null;
			try {
				this.$router.push({ path: "/" }); // reindirizza al login
			} catch (e) {
				this.errormsg = e.toString(); // gestisco errori
			}
			this.loading = false;
		},
		async getChat() { // funzione che prende tutte le chat dal databse
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/" + this.id + "/conversations", { // chiama la query che trova le chat
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"), // passa il token alla query tramite json
					}

				});
				this.chats = response.data; // i dati in risposta della query
				this.createDynamicListsFromJSON(this.chats); // chiama la funzione
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		createDynamicListsFromJSON(json) { // funzione che crea dinamicamente la lista delle chat
			this.dynamicData = json; // assegnazione dati
			if (this.dynamicData) {
				this.dynamicData.forEach((item, index) => { // per ogni sezione groupPhoto specifica che si tratti di un immagine
					item.GroupPhoto = "data:image/jpeg;base64," + item.GroupPhoto;
				});
			}

		},
		closePopup() { // chiude tutti i popup
			this.showPopup = false;
			this.showPopup1 = false;
		},

		async creaGruppo() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/users/" + this.id + "/conversations", {
					sgn_name: this.newGroupname,
				},
					{
						headers: {
							Authorization: "Bearer " + localStorage.getItem("token"),
						}
					}); // crea un json che gli passa un nome
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.getChat()
		},

		async creaConversazione() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/users/" + this.id, {
					sgn_name: this.newUser,
				},
					{
						headers: {
							Authorization: "Bearer " + localStorage.getItem("token"),
						}
					}); // crea un json che gli passa un nome
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.getChat()
		},

		async apriChat(conv) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/" + this.id + "/conversations/" + conv.IdChat + "/messages", { // chiama la query che trova le chat
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"), // passa il token alla query tramite json
					}

				});
				this.messages = response.data; // i dati in risposta della query
				this.createDynamicListsFromJSON_2(this.messages); // chiama la funzione
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		createDynamicListsFromJSON_2(json) { // funzione che crea dinamicamente la lista dei messaggi
			this.dynamicData2 = json; // assegnazione dati
			if (this.dynamicData2) {
				this.dynamicData2.forEach((item, index) => { // per ogni sezione groupPhoto specifica che si tratti di un immagine
					item.Foto = "data:image/jpeg;base64," + item.Foto;
				});
			}

		},
		conversioneUnix(tempo){
			const date = new Date(tempo);
			return date.toLocaleString();
		},

		async getNome(identifier) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/" + identifier, { // chiama la query che trova il nome
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"), // passa il token alla query tramite json
					}

				});
				this.nomi = response.data; // i dati in risposta della query
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			return this.nomi
		},

	},
	mounted() {
		this.getChat(); // per chiamare le funzioni istantaneamente al caricamento dalla pagnia

	}
}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<!-- Lateral Section -->
			<nav class="col-md-3 col-lg-2 d-md-block bg-light sidebar">
				<div class="sidebar-sticky">
					<!-- Sidebar Title -->
					<div class="sidebar-title" style="text-align: center;">
						<h5 class="me-2">Chat</h5>
						<div class="btn-group me-2">

							<button type="button" class="btn btn-sm btn-outline-primary" @click="showPopup = true">
								create group
							</button>
							<button type="button" class="btn btn-sm btn-outline-primary" @click="showPopup1 = true">
								new chat
							</button>
						</div>
					</div>

					<!-- Scrollable List -->
					<div class="list-container">
						<ul>
							<li v-for="(item, index) in chats" :key="index" @click="apriChat(item)">
								<a class="clickable-item">
									<img :src="item.GroupPhoto" alt="" class="group-photo">
									{{ item.GroupName }}
								</a>
							</li>
						</ul>
					</div>
				</div>
			</nav>


			<!-- Body Section -->
			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<div
					class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
					<h1 class="h2">Hello, {{ username }}!</h1>
					<div class="btn-toolbar mb-2 mb-md-0">
						<div class="btn-group me-2">
							<button type="button" class="btn btn-sm btn-outline-primary" @click="newPage">
								Profile
							</button>
							<button type="button" class="btn btn-sm btn-outline-primary" @click="login">
								Disconnect
							</button>
						</div>
					</div>
				</div>

				<!-- Body Content -->
				<div class="body-content" style="overflow-y: auto; max-height: 75vh;">
					<p>Welcome to the main content area. Here is where you can display information.</p>

					<div class="chatlist-container">
						<div class="list-container">
							<ul>
								<li v-for="(item, index) in messages" :key="index"
									:class="{ 'user-message': item.User == this.id }">
									<a class="chatclickable-item">
										<!-- non funziona la funzione getnome -->
										<span class="user">{{ getNome(item.User) }}</span>
										<span ><img :src="item.Foto" alt=" "> </span>
										<span class="user"></span>
										<span class="chatmessage-text">{{ item.Messaggio }}</span>
										<div class="chattimestamp">
											{{ conversioneUnix(item.Timestamp* 1000) }}
											<span class="chatcheckmarks">{{ item.Checkmarks }}</span>
										</div>
									</a>
								</li>
							</ul>
						</div>
					</div>

				</div>

				<!-- Bottom Section -->
				<div class="bottom-section mt-4 pt-3 border-top">
					<h5>Footer Section</h5>
					<p>This is the bottom section of the body.</p>
				</div>
			</main>
		</div>

		<!-- Error Message -->
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
	<div v-if="showPopup" class="popup-overlay" @click.self="closePopup">
		<div class="popup-content">
			<h2>New group name</h2>
			<input type="text" v-model="newGroupname" placeholder="Type here..." />
			<div class="popup-actions">
				<button @click="creaGruppo">Submit</button>
				<button @click="closePopup">Close</button>
			</div>
		</div>
	</div>
	<div v-if="showPopup1" class="popup-overlay" @click.self="closePopup">
		<div class="popup-content">
			<h2>User to contact</h2>
			<input type="text" v-model="newUser" placeholder="Type here..." />
			<div class="popup-actions">
				<button @click="creaConversazione">Submit</button>
				<button @click="closePopup">Close</button>
			</div>
		</div>
	</div>
</template>

<style>
/* Sidebar styling */
.sidebar {
	height: 100vh;
	position: sticky;
	top: 0;
	overflow: hidden;
	/* Prevent overflow of the entire sidebar */
	display: flex;
	flex-direction: column;
}

.sidebar-title {
	padding: 10px 20px;
	border-bottom: 1px solid #ddd;
	background-color: #f8f9fa;
	font-weight: bold;
	color: #333;
}

.sidebar-scroll {
	overflow-y: auto;
	/* Add scroll for the list */
	flex-grow: 1;
	/* Allow the list to grow and take remaining height */
	padding: 10px;
}

.sidebar ul {
	list-style: none;
	padding: 0;
}

.sidebar ul li {
	margin-bottom: 10px;
}

.sidebar ul li a {
	display: block;
	padding: 10px;
	text-decoration: none;
	color: #333;
	border-radius: 4px;
	transition: background-color 0.3s ease;
}

.sidebar ul li a:hover {
	background-color: #007bff;
	color: white;
}

.sidebar ul li a.active {
	background-color: #0056b3;
	color: white;
}

/* Scrollbar customization (optional) */
.sidebar-scroll::-webkit-scrollbar {
	width: 8px;
}

.sidebar-scroll::-webkit-scrollbar-thumb {
	background-color: #007bff;
	border-radius: 4px;
}

.sidebar-scroll::-webkit-scrollbar-thumb:hover {
	background-color: #0056b3;
}


/* Main content styling */
.body-content {
	padding: 20px;
	background-color: #f9f9f9;
	border-radius: 8px;
	box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.bottom-section {
	background-color: #f1f1f1;
	padding: 10px 20px;
	border-radius: 8px;
}

.chat-list {
	padding: 0;
	list-style: none;
	margin: 0;
}

.chat-item {
	display: flex;
	align-items: center;
	padding: 10px;
	border-bottom: 1px solid #ddd;
	cursor: pointer;
	transition: background-color 0.2s;
}

.chat-item:hover {
	background-color: #f0f0f0;
}

.group-photo {
	width: 50px;
	height: 50px;
	border-radius: 50%;
	margin-right: 15px;
}

.chat-details {
	display: flex;
	flex-direction: column;
}

.group-name {
	font-size: 16px;
	font-weight: bold;
	margin: 0;
}

.group-description {
	font-size: 14px;
	color: #666;
	margin: 0;
}

.chatbody-content {
	background-color: #ffffff;
	width: 400px;
	height: 600px;
	border-radius: 10px;
	box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
	overflow-y: auto;
	display: flex;
	flex-direction: column;
}

.chatlist-container {
	padding: 10px;
	flex-grow: 1;
	display: flex;
	flex-direction: column;
}

ul {
	list-style-type: none;
	padding: 0;
	margin: 0;
}

li {
	margin-bottom: 10px;
	display: flex;
	align-items: flex-end;
}

.chatclickable-item {
	display: block;
	max-width: 80%;
	padding: 10px;
	border-radius: 10px;
	background-color: #dcf8c6;
	color: #333;
	font-size: 14px;
	text-decoration: none;
	word-wrap: break-word;
	position: relative;
	box-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.chatclickable-item img {
	width: 200px;
	height: 200px;
	margin-right: 8px;
}



.chattimestamp {
	font-size: 12px;
	color: #999;
	margin-top: 5px;
	text-align: right;
}


.chatclickable-item .user {
	font-weight: bold;
	display: block;
	margin-bottom: 5px;
}



.clickable-item .message-text {
	margin-bottom: 5px;
}

.user-message {
  display: flex;
 
  justify-content: flex-end; /* Aligns the content to the right */
  text-align: right; /* Ensures the text aligns to the right */
  margin: 5px 10px; /* Adds some spacing around the message */
}


</style>
