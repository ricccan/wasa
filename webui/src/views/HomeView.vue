<script>
export default {
	data: function () {
		return {
			componentKey: 0, // Key to force re-render
			errormsg: null,
			loading: false,
			some_data: null,
			token: localStorage.getItem("token"), // prende il valore del token dato dal login
			username: this.$route.query.username, // prende il valore dell'username dato nel login
			chats: {}, // json che conterrà la lista delle chat
			id:  this.$route.query.id, // id utente collegato allo username
			dynamicData: {}, // json che utilizzo insieme a chats per gestire le conversazioni
			showPopup: false, // popup che apre la sezione "crea gruppo"
			showPopup1: false, // popup che apre la sezione "nuova chat"
			showPopup2: false, // popup che apre la sezione info gruppo
			newGroupname: null, // valore che contiene il nome da dare al gruppo che sta venendo creato
			newUser: null, // valore dello username dell'utente con cui si vuole creare una nuova conversazione
			messages: {}, // json a cui passo i messaggi
			dynamicData2: {}, // json che utilizzo insieme a chats per gestire i messaggi
			nomi: {}, // variabile a cui do il valore dei nomi nei messaggi
			selectedFile: null, // file messo da input
			imagePreview: null, // preview
			addUser: null, // variabile per utente da aggiungere nel gruppo
			removeUser: null, // variabile per utente da rimuovere dal gruppo
			newName: null, // variabile per nuovo nome del gruppo
			newId: {}, // id che prendo dalla funzione getID
			idGroup: null, // id del gruppo
			changeGroupPhoto: null, // valore che contiene la foto da dare al gruppo che sta venendo creato
			newMessage: null, // valore che contiene il messaggio che si vuole inviare
			currentChat: null, // chat attuale
			entrato: false, // variabile che mi dice se sono entrato in una chat
			showPopup3: false, // popup che apre la sezione info messaggio
			idTemp: null, // id temporaneo
			tempMessId: null, // id temporaneo messaggio
			forwardedUser: null, // utente a cui inoltrare il messaggio
			showPopup4: false, // popup che apre la sezione inoltro messaggio
			showPopup5: false, // popup che apre la sezione risposta messaggio
			respondMessage: null, // messaggio di risposta
			oneMessage: {}, // messaggio singolo
			datiTemp: {}, // dati temporanei
			showPopup6: false,
			reazione: null,
			showPopup7: false,
			reactions: {}, // json a cui passo le reazioni
			showPopup8: false,
			forwardName: null, // nome di chi inoltro il messaggio
			userList: {}, // lista utenti
			chatCreated: null,
			groupUsers: {}, // utenti del gruppo
			showPopup9: false,
			primaChat: null,
			showPopup10: false,
			chatInterval: null, // Store interval reference
			
		}
	},
	methods: {
		unselectFile() {
			this.respondMessage = null;
			this.imagePreview = null;
			this.newMessage = null;
			this.selectedFile = null;
			let fileInput = document.getElementById('photoInput');
			if (fileInput) {
				fileInput.value = '';
			}
			fileInput = document.getElementById('photoInput2');
			if (fileInput) {
				fileInput.value = '';
			}
			fileInput = document.getElementById('photoInput3');
			if (fileInput) {
				fileInput.value = '';
			}
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
		async newPage() { // funzione che porta alla pagina del profilo
			this.loading = true;
			this.errormsg = null;
			try {
				this.$router.push({ 
					path: "/profile",
					query: { username: this.username, id: this.id },
				 }); // reindirizza alla pagina del profilo
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
		async getChat() { // getMyConversations (prende tutte le conversazioni dell'utente)
			this.loading = true;
			
			try {
				let response = await this.$axios.get("/users/" + this.id + "/conversations", { // chiama la query che trova le chat
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"), // passa il token alla query tramite json
					}

				});
				this.chats = response.data; // i dati in risposta della query
				if (this.chats) this.chats = this.chats.sort((a, b) => new Date(b.LastChange) - new Date(a.LastChange)); // ordina le chat per data
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
			this.errormsg=null;
			this.showPopup = false;
			this.showPopup1 = false;
			this.showPopup2 = false;
			this.showPopup3 = false;
			this.showPopup4 = false;
			this.showPopup5 = false;
			this.showPopup6 = false;
			this.showPopup7 = false;
			this.showPopup8 = false;
			this.showPopup9 = false;
			this.showPopup10 = false;
		},

		async creaGruppo() { // createGroup (crea un gruppo)
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
			this.closePopup()
		},

		async creaConversazione() { // createChat (crea una chat)
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
					this.chatCreated = response.data;
					if (response.data == 0){
						this.errormsg = "user does not exist";
						return
					}
			} catch (e) {
				this.errormsg = "user does not exist";
				var temp = 1;
				
			}
			this.loading = false;
			
			if (this.chatCreated != 0){
				this.getChat()
				this.showPopup10 = true;
			}
			
			return this.chatCreated 
		},

		async apriChat(conv) { // getMessages + setSeen (entra in chat, fa uscire i messaggi e li segna come letti)
			this.loading = true;
			

			if (this.chatInterval) {
        	clearInterval(this.chatInterval);
    		}

			// Clear any previous interval to prevent multiple instances
			if (this.chatInterval) {
            clearInterval(this.chatInterval);
        	}


			try {
				let response1 = await this.$axios.post("/users/" + this.id + "/conversations/" + conv, { // chiama la query che setta i visualizzati
				});
			} catch (e) {
				this.errormsg = e.toString();
			}
			// this.loading = false;
			try {
				let response = await this.$axios.get("/users/" + this.id + "/conversations/" + conv + "/messages", { // chiama la query che trova le chat
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
			this.getChat();
			this.chatInterval = setInterval(() => {
        	this.apriChat(conv);
    		}, 1000);
			
			// Set an interval to refresh messages every 5 seconds
		},
		createDynamicListsFromJSON_2(json) { // funzione che crea dinamicamente la lista dei messaggi
			this.dynamicData2 = json; // assegnazione dati
			if (this.dynamicData2) {
				this.dynamicData2.forEach((item, index) => { // per ogni sezione groupPhoto specifica che si tratti di un immagine
					item.Foto = "data:image/jpeg;base64," + item.Foto;
				});
			}

		},
		conversioneUnix(tempo) {
			const date = new Date(tempo);
			return date.toLocaleString();
		},


		async getId(nome) { // getUserId (ritorna l'id di uno user dato il nome utente)
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/usernames/" + nome, {
					headers: { // Headers should be part of the same object
						Authorization: "Bearer " + localStorage.getItem("token"),
					}
				}); // crea un json che gli passa un nome
				this.newId = response.data; // i dati in risposta della query
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			return this.newId
		},

		async addMember() { // addToGroup (aggiunge membri al gruppo)
			this.loading = true;
			this.errormsg = null;
			var temp = await this.getId(this.addUser)
			if (temp){
				try {
				let response = await this.$axios.post("/users/" + this.id + "/groups", {
					adg_utente_id: temp, // forse giusto?
					adg_group_id: this.idGroup // FORSE GIUSTO?
				},
					{
						headers: {
							Authorization: "Bearer " + localStorage.getItem("token"),
						}
					}); // crea un json che gli passa un nome
			} catch (e) {
				this.errormsg = e.toString();
				
			} 
			} else {
				this.errormsg = "user doesnt exist"
			}
			
			this.showPopup8 = true;
			this.loading = false;
			
		},

		async removeMember() { // removeFromGroup (elimina un membro del gruppo)
			this.loading = true;
			this.errormsg = null;
			var temp = await this.getId(this.removeUser)
			if (temp == this.id){
				this.errormsg = "you can't remove yourself"
				this.loading = false;
				return
			}
			try {
				await this.$axios.delete("/users/" + temp + "/groups/" + this.idGroup, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
					}
				}); // crea un json che gli passa un nome
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.closePopup()
			this.getChat()
		},
		async removeSelf() { // removeFromGroup (elimina se stesso dal gruppo)
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/users/" + this.id + "/groups/" + this.idGroup, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
					}
				}); // crea un json che gli passa un nome
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.closePopup()
			this.getChat()
			this.entrato = false;
		},

		async cambiaNomeGruppo() { // setGroupName (cambia il nome del gruppo)
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/users/" + this.id + "/conversations/" + this.idGroup + "/groupName", {
					sgn_name: this.newName,
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
			this.getChat();
			this.closePopup()
		},

		async cambiaFoto() { // setGroupPhoto (cambia foto profilo del gruppo)
			this.loading = true;
			this.errormsg = null;
			try {
				let formData = new FormData();
				formData.append("photo", this.selectedFile); // forse da cambiare perchè sovrascrive la foto profilo

				let response = await this.$axios.post("/users/" + this.id + "/conversations/" + this.idGroup + "/groupphoto", formData, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
						"Content-Type": "multipart/form-data",
					}
				});

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.getChat();
			this.closePopup()
		},

		async sendMessage() { // sendMessage (manda i messaggi)
			this.loading = true;
			this.errormsg = null;
			try {
				let formData2 = new FormData();
				if (this.newMessage == null) {
					this.newMessage = " ";
				}
				formData2.append("photo", this.selectedFile); // forse da cambiare perchè sovrascrive la foto profilo
				formData2.append("messageText", this.newMessage);
				if (this.selectedFile) {
					formData2.append("IsPhoto", "true");
				} else {
					formData2.append("IsPhoto", "false");
				}
				await this.$axios.post("/users/" + this.id + "/conversations/" + this.idGroup + "/messages", formData2, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
						"Content-Type": "multipart/form-data",
					}
				});

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.apriChat(this.idGroup)
			this.unselectFile() // da usare anche dalle altre parti
			this.getChat()
		},
		async deleteMessage() { // deleteMessage (cancella un messaggio)
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/users/" + this.id + "/conversations/" + this.idGroup + "/messages/" + this.tempMessId, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
					}
				}); // crea un json che gli passa un nome
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.closePopup()
			this.apriChat(this.idGroup)
			this.getChat()
		},
		async forward(messi) { // forwardMessage (inoltra il messaggio)
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/users/" + this.id + "/conversations/" + this.idGroup + "/messages/" + this.tempMessId, {
					chatId: messi,
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
			this.getChat();
			this.closePopup();
			this.apriChat(messi)
			this.getChat()
		},
		async respond(risponde) { // commentMessage (risponde ad un messaggio)
			this.loading = true;
			this.errormsg = null;
			try {
				let formData2 = new FormData();

				if (this.respondMessage == null) {
					this.respondMessage = " ";
				}

				formData2.append("photo", this.selectedFile); // forse da cambiare perchè sovrascrive la foto profilo
				formData2.append("messageText", this.respondMessage);
				if (this.selectedFile) {
					formData2.append("IsPhoto", "true");
				} else {
					formData2.append("IsPhoto", "false");
				}
				await this.$axios.post("/users/" + this.id + "/conversations/" + this.idGroup + "/messages/" + risponde + "/comments", formData2, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
						"Content-Type": "multipart/form-data",
					}
				});

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.apriChat(this.idGroup)
			this.closePopup()
		},
		async react(emoji) { // reactMessage (aggiunge una reazione)
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/users/" + this.id + "/conversations/" + this.idGroup + "/messages/" + this.tempMessId + "/reactions", {
					reaction: emoji,
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
			this.getChat();
			this.closePopup();

		},
		async deleteReaction() { // uncommentMessage (cancella la reazione)
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/users/" + this.id + "/conversations/" + this.idGroup + "/messages/" + this.tempMessId + "/reactions", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
					}
				}); // crea un json che gli passa un nome
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.closePopup()
			this.apriChat(this.idGroup)
		},
		async getReaction(conv, messId) { // getReactions (prende le reazioni)
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/" + this.id + "/conversations/" + conv + "/messages/" + messId + "/reactions", { // chiama la query che trova le chat
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"), // passa il token alla query tramite json
					}

				});
				this.reactions = response.data; // i dati in risposta della query


			} catch (e) {
				this.errormsg = e.toString();

			}
			this.loading = false;
			

		},
		async createForward(){
			var nuovoid = await this.creaConversazione()
			if (nuovoid != null){
				
				this.forward(nuovoid)
			} else {
				this.errormsg = "user does not exist"
			}
			
				
		},
		async getPartecipants(){ // getConversation (prende la lista dei partecipanti di un gruppo)
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/" + this.id + "/conversations/" + this.idGroup, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
					}
					
				});
				this.groupUsers = response.data; // crea un json che gli passa un nome
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.showPopup9 = true
		},
		
	},
	mounted() {
		
		this.getChat(); // per chiamare le funzioni istantaneamente al caricamento dalla pagnia
		this.chatInterval = setInterval(() => {
      this.getChat();
    }, 1000); 
		
	},
	beforeUnmount() {
    clearInterval(this.chatInterval); // Cleanup when component is destroyed
  },
	
}
</script>

<template>
	
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css">
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
								Create group
							</button>
							<button type="button" class="btn btn-sm btn-outline-primary" @click="showPopup1 = true">
								New chat
							</button>
						</div>
					</div>

					<!-- Scrollable List -->
					<div class="list-container">
						<ul style="list-style: none; padding: 0; margin: 0;">
							<li v-for="(item, index) in chats" :key="index"
								style="display: flex; align-items: center; padding: 10px; border-bottom: 1px solid #e6e6e6; cursor: pointer;"
								@click="apriChat(item.IdChat), idGroup = item.IdChat, currentChat = item.GroupName, entrato = true">
								<a
									style="display: flex; align-items: center; width: 100%; text-decoration: none; color: inherit;">
									<img :src="item.GroupPhoto"
										style="width: 50px; height: 50px; border-radius: 50%; object-fit: cover; margin-right: 15px;">
									<div style="flex-grow: 1;">
										<p style="margin: 0; font-size: 16px; font-weight: bold; color: #333;">
											{{ item.GroupName }}
										</p>
										<p style="margin: 5px 0 0; font-size: 14px; color: #666;">
											{{ item.Snippet }}
										</p>
									</div>
									<button v-if="item.Group"
										style="background: none; border: none; cursor: pointer; padding: 5px; color: #666; transition: color 0.2s;"
										@click.stop="showPopup2 = true, idGroup = item.IdChat">
										<i class="fas fa-pencil-alt" style="font-size: 16px; color: black;"></i>
									</button>
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
				<div class="body-content" v-if="entrato" style="overflow-y: auto; max-height: 60vh;">

					<div class="chatlist-container">
						<div class="list-container">
							<ul>
								<li v-for="(item, index) in messages" :key="index"
									:class="{ 'user-message': item.User == this.id }">
									<div class="chatclickable-item">
										<span class="user" v-if="item.Inoltrato"> forwarded -> </span>
										<div class="response-row">
											<!-- da sistemare il css della risposta-->
											<div class="response-row" v-if="item.Risposta"
												style="background-color: #f0f0f0; border-left: 4px solid #34b7f1; padding: 5px 10px; margin-bottom: 8px; border-radius: 5px;">
												<span class="response-text">
													{{ messages.find(msg => msg.IdMess == item.Risposta)?.Messaggio }}
												</span>
											</div>

										</div>
										<div style="display: flex; align-items: center;">
											<button class="btn btn-outline-light mb-2"
												@click="idTemp = item.User, showPopup3 = true, tempMessId = item.IdMess"
												style="margin-right: 10px;">
												<i class="fas fa-pencil-alt" style="color: black;"></i>
											</button>
											<span class="user">{{ item.Username }}</span>
											<button class="btn btn-outline-light mb-2"
												@click="idTemp = item.User, showPopup7 = true, tempMessId = item.IdMess, getReaction(idGroup, item.IdMess)"
												style="margin-right: 10px;">
												<i class="fas fa-eye" style="color: black;"></i>
											</button>
										</div>

										<span><img :src="item.Foto" alt=" "> </span>
										<span class="user"></span>
										<span class="chatmessage-text">{{ item.Messaggio }}</span>
										<div class="chattimestamp">
											{{ conversioneUnix(item.Timestamp * 1000) }}
											<span v-if="item.Checkmarks == 0 && item.User == this.id"
												:style="{ color: item.Checkmarks === 1 ? '#53bdeb' : 'gray', fontSize: '16px' }">
												&#xe013;
											</span>
											<span v-if="item.Checkmarks >= 1 && item.User == this.id"
												:style="{ color: item.Checkmarks === 1 ? '#53bdeb' : 'gray', fontSize: '16px' }">
												&#xe013;
											</span>
											<span v-if="item.Checkmarks >= 1 && item.User == this.id"
												:style="{ color: item.Checkmarks === 1 ? '#53bdeb' : 'gray', fontSize: '16px' }">
												&#xe013;
											</span>
										</div>

									</div>
								</li>
							</ul>
						</div>
					</div>

				</div>
				<!-- Bottom Section -->
				<div class="bottom-section mt-4 pt-3 border-top" v-if="entrato">
					<form class="footer-form text-center" method="post" enctype="multipart/form-data">
						<div class="form-group">
							<label for="messageInput">Your Message:</label>
							<input type="text" id="messageInput" v-model="newMessage" class="form-control"
								name="message" placeholder="Enter your message" required>
						</div>
						<div class="form-group d-flex align-items-center mt-2">
							<label for="photoInput" class="me-2">Attach a Photo:</label>
							<input type="file" id="photoInput" class="form-control me-2" @change="onFileChange"
								name="photo" accept="image/*">
							<button @click="unselectFile" class="btn btn-danger">
								<i class="fas fa-trash-alt"></i>
							</button>
						</div>

						<button v-if="newMessage || selectedFile" @click.prevent="sendMessage" type="submit"
							class="btn btn-primary mt-3">Send to {{
								currentChat }}</button>
					</form>
				</div>
			</main>
		</div>

		<!-- Error Message -->
		
	</div>
	<div v-if="showPopup" class="popup-overlay" @click.self="closePopup">
		<div class="popup-content">
			<h2>New group name</h2>
			<input type="text" v-model="newGroupname" placeholder="Type here..." />
			<div class="popup-actions">
				<button v-if="newGroupname" @click="creaGruppo">Submit</button>
				<button @click="closePopup">Close</button>
			</div>
		</div>
	</div>
	<div v-if="showPopup1" class="popup-overlay" @click.self="closePopup">
		<div class="popup-content">
			<h2>User to contact</h2>
			<input type="text" v-model="newUser" placeholder="Type here..." />
			<div class="popup-actions">
				<button v-if="newUser" @click="creaConversazione">Submit</button>
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<button @click="closePopup">Close</button>
			</div>
		</div>
	</div>
	<div v-if="showPopup2" class="popup-overlay" @click.self="closePopup">
		<!-- TODO: cambiare i pulsanti-->
		<div class="popup-content">
			<h2>Group action</h2>
			<a>Add member</a>
			<div style="display: flex; align-items: center;">
				<input type="text" v-model="addUser" placeholder="Type here..." style="margin-right: 10px;" />
				<button @click="addMember" class="btn btn-primary">Submit</button>
			</div>
			<a>Remove member</a>
			<div style="display: flex; align-items: center;">
				<input type="text" v-model="removeUser" placeholder="Type here..." style="margin-right: 10px;" />
				<button @click="removeMember" class="btn btn-primary">Submit</button>
			</div>
			<a>Change name</a>
			<div style="display: flex; align-items: center;">
				<input type="text" v-model="newName" placeholder="Type here..." style="margin-right: 10px;" />
				<button v-if="newName" @click="cambiaNomeGruppo" class="btn btn-primary">Submit</button>
			</div>
			<a>Change photo</a>
			<div style="display: flex; align-items: center;">
				<input type="file" id="photoInput2" @change="onFileChange" accept="image/*"
					style="margin-right: 10px;" />
				<button @click="cambiaFoto" class="btn btn-primary">Submit</button>
			</div>
			<div v-if="imagePreview" class="image-preview">
				<img :src="imagePreview" alt="Selected photo preview" />
			</div>
			<button v-if="imagePreview" @click="unselectFile" class="btn btn-danger">
				<i class="fas fa-trash-alt"></i>
			</button>
			<div class="popup-actions">
				<button  class="btn btn-danger" @click="getPartecipants">See partecipants</button> <!-- da implementare -->
				<button  class="btn btn-danger" @click="removeSelf">Exit group</button>
			</div>
			<div class="popup-actions">
				<button @click="closePopup">Close</button>
			</div>
			<div class="alert alert-danger" role="alert" v-if="errormsg">
				{{errormsg}}
			</div>
		</div>
	</div>
	<div v-if="showPopup3" class="popup-overlay" @click.self="closePopup">
		<!-- TODO: Update button styles if needed -->
		<div class="popup-content">
			<h2>Message Actions</h2>

			<!-- Delete Message Section -->
			<div v-if="idTemp == this.id"
				style="display: flex; align-items: center; justify-content: center; margin-bottom: 1rem;">
				<div style="text-align: center; margin-bottom: 1rem;">
					<div>Delete message</div>
					<button v-if="idTemp == this.id" @click="deleteMessage" class="btn btn-danger">
						<i class="fas fa-trash-alt"></i>
					</button>
				</div>
			</div>

			<!-- First Row: Forward and React to Message -->
			<div style="display: flex; justify-content: space-around; margin-bottom: 1rem;">
				<div style="text-align: center;">
					<div>Forward message</div>
					<button @click="showPopup4 = true" class="btn btn-success">
						<i class="fas fa-paper-plane"></i>
					</button>
				</div>
				<div style="text-align: center;">
					<div>React to message</div>
					<button @click="showPopup6 = true" class="btn btn-primary">
						<i class="fas fa-spell-check"></i>
					</button>
				</div>
			</div>

			<!-- Second Row: Reply and Delete Reaction -->
			<div style="display: flex; justify-content: space-around; margin-bottom: 1rem;">
				<div style="text-align: center;">
					<div>Reply to message</div>
					<button @click="showPopup5 = true" class="btn btn-success">
						<i class="fas fa-reply"></i>
					</button>
				</div>
				<div style="text-align: center;">
					<div>Delete reaction</div>
					<button @click="deleteReaction" class="btn btn-danger">
						<i class="fas fa-remove-format"></i>
					</button>
				</div>
			</div>

			<!-- Close Button -->
			<div class="popup-actions">
				<button @click="closePopup">Close</button>
			</div>
		</div>
	</div>

	<div v-if="showPopup4" class="popup-overlay" @click.self="closePopup">
		<!-- TODO: cambiare i pulsanti -->
		<div class="popup-content">
			<h2>Forward</h2>
			
			<div class="list-container"
				style="max-height: 300px; overflow-y: auto; border: 1px solid #ccc; padding: 1rem; border-radius: 8px; background-color: #f9f9f9;">
				<div style="display: flex; align-items: center;">
				<input type="text" v-model="newUser" placeholder="username..." style="margin-right: 10px;" />
				<button @click="createForward" class="btn btn-primary">Send</button>
			</div>
				<ul style="list-style-type: none; margin: 0; padding: 0;">
					<li v-for="(item, index) in chats" :key="index"
						style="display: flex; align-items: center; margin-bottom: 1rem; border-bottom: 1px solid #ddd; padding-bottom: 0.5rem;">
						<img :src="item.GroupPhoto" alt="" class="group-photo"
							style="width: 40px; height: 40px; border-radius: 50%; margin-right: 1rem;">
						<div style="flex-grow: 1;">
							<span>{{ item.GroupName }}</span>
						</div>
						
						<button class="edit-button mb-2" @click="forward(item.IdChat)"
							style="padding: 0.5rem 1rem; background-color: #007bff; color: #fff; border: none; border-radius: 5px; cursor: pointer;">Send</button>
					</li>
				</ul>

				

			</div>

			<div class="popup-actions"
				style="display: flex; justify-content: center; align-items: center; gap: 1rem; margin-top: 1rem;">
				<button @click="showPopup4 = false" type="submit" class="btn btn-primary">
					<i class="fas fa-chevron-left"></i>
				</button>
				<div>
					<div v-if="errormsg" class="alert alert-danger" role="alert">
						User does not exist
					</div>
				</div>
				<button @click="closePopup"
					style="background-color: #dc3545; color: #fff; border: none; border-radius: 5px; cursor: pointer;">
					Close
				</button>
			</div>


		</div>
	</div>

	<div v-if="showPopup5" class="popup-overlay" @click.self="closePopup">
		<div class="bottom-section mt-4 pt-3 border-top">
			<form class="footer-form text-center" method="post" enctype="multipart/form-data">
				<div class="form-group">
					<label for="messageInput">Your Message:</label>
					<input type="text" id="messageInput" v-model="respondMessage" class="form-control" name="message"
						placeholder="Enter your message" required>
				</div>
				<div class="form-group mt-2">
					<label for="photoInput3">Attach a Photo:</label>
					<input type="file" id="photoInput3" class="form-control" @change="onFileChange" name="photo"
						accept="image/*">
					<div v-if="imagePreview" class="image-preview">
						<img :src="imagePreview" alt="Selected photo preview" />
					</div>
					<button v-if="imagePreview" @click="unselectFile" class="btn btn-danger">
						<i class="fas fa-trash-alt"></i>
					</button>
				</div>
				<div style="display: flex; gap: 1rem; justify-content: flex-start;">
					<button @click="showPopup5 = false" type="submit" class="btn btn-primary mt-3">
						<i class="fas fa-chevron-left"></i>
					</button>
					<button v-if="respondMessage || selectedFile" @click.prevent="respond(tempMessId)" type="submit"
						class="btn btn-success mt-3">
						Send to {{ currentChat }}
					</button>
				</div>



			</form>
		</div>
	</div>
	<div v-if="showPopup6" class="popup-overlay" @click.self="closePopup">
		<div class="bottom-section mt-4 pt-3 border-top">
			<form class="footer-form text-center" method="post" enctype="multipart/form-data">
				<div class="emoji-buttons">
					<button type="button" @click="react('&#128511;')" class="emoji-button">&#128511;</button>
					<button type="button" @click="react('&#128512;')" class="emoji-button">&#128512;</button>
					<button type="button" @click="react('&#128513;')" class="emoji-button">&#128513;</button>
					<button type="button" @click="react('&#128514;')" class="emoji-button">&#128514;</button>
					<button type="button" @click="react('&#128517;')" class="emoji-button"> &#128517;</button>
				</div>
				<div class="popup-header" style="display: flex; justify-content: space-between; align-items: center;">
					<button @click="showPopup6 = false" type="submit" class="btn btn-primary mt-3">
						<i class="fas fa-chevron-left"></i>
					</button>
					<div class="popup-actions">
						<button type="button" @click="closePopup" class="close-button">Close</button>
					</div>
				</div>

			</form>
		</div>
	</div>
	<div v-if="showPopup7" class="popup-overlay" @click.self="closePopup">
		<div class="bottom-section mt-4 pt-3 border-top">
			<h1>Reactions</h1>
			<form class="footer-form text-center" method="post" enctype="multipart/form-data">
				<div class="reaction-list-container" style="display: flex; flex-direction: column; height: 300px;">
					<!-- Scrollable List -->
					<ul class="reaction-list"
						style="overflow-y: auto; height: 100%; padding: 0; margin: 0; list-style: none;">
						<li v-for="(item, index) in reactions" :key="index"
							style="padding: 10px; border-bottom: 1px solid #ddd;">
							<span class="chatmessage-text">{{ item.User }}</span>
							<span class="chatmessage-text" style="margin-left: 10px;">{{ item.Reazione }}</span>
						</li>
					</ul>

					<!-- Popup Header with Buttons -->
					<div class="popup-header"
						style="display: flex; justify-content: space-between; align-items: center; margin-top: 10px;">
						<button @click="showPopup7 = false" type="submit" class="btn btn-primary">
							<i class="fas fa-chevron-left"></i>
						</button>
						<div class="popup-actions">
							<button type="button" @click="closePopup" class="close-button">Close</button>
						</div>
					</div>
				</div>

			</form>
		</div>
	</div>
	<div v-if="showPopup8" class="popup-overlay" @click.self="closePopup">
		<div class="bottom-section mt-4 pt-3 border-top">
			<h1 v-if="!errormsg">User successfully added</h1>
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
						<div class="popup-actions">
							<button type="button" @click="closePopup" class="close-button">Close</button>
						</div>
			
		</div>		
	</div>
	<div v-if="showPopup9" class="popup-overlay" @click.self="closePopup">
		<div class="bottom-section mt-4 pt-3 border-top">
			<h1>Users in the group</h1>
			<form class="footer-form text-center" method="post" enctype="multipart/form-data">
				<div class="reaction-list-container" style="display: flex; flex-direction: column; height: 300px;">
					<!-- Scrollable List -->
					<ul class="reaction-list"
						style="overflow-y: auto; height: 100%; padding: 0; margin: 0; list-style: none;">
						<li v-for="(item, index) in groupUsers" :key="index"
							style="padding: 10px; border-bottom: 1px solid #ddd;">
							<span class="chatmessage-text">{{ item.u_username }}</span>
						</li>
					</ul>

					<!-- Popup Header with Buttons -->
					<div class="popup-header"
						style="display: flex; justify-content: space-between; align-items: center; margin-top: 10px;">
						<div class="popup-actions">
							<button type="button" @click="closePopup" class="close-button">Close</button>
						</div>
					</div>
				</div>

			</form>
		</div>
	</div>
	<div v-if="showPopup10" class="popup-overlay" @click.self="closePopup">
		<div class="bottom-section mt-4 pt-3 border-top">
			<h1>New chat created</h1>
			<div class="popup-actions">
							<button type="button" @click="closePopup" class="close-button">Close</button>
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
	/* Sidebar title styling */
	padding: 10px 20px;
	border-bottom: 1px solid #ddd;
	background-color: #f8f9fa;
	font-weight: bold;
	color: #333;
}

.sidebar ul li {
	/* Sidebar list item styling */
	margin-bottom: 10px;
	border-bottom: 1px solid #ddd;
	/* Adds a border between items */
}

.sidebar ul li {
	position: relative;
}

.sidebar ul li a {
	display: block;
	padding: 20px;
	text-decoration: none;
	color: #333;
	border-radius: 4px;
	transition: background-color 0.3s ease;
	width: 100%;
	height: 100%;
	box-sizing: border-box;
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
	/* Chat item styling */
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

	justify-content: flex-end;
	/* Aligns the content to the right */
	text-align: right;
	/* Ensures the text aligns to the right */
	margin: 5px 10px;
	/* Adds some spacing around the message */
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

.response-row {
	display: flex;
	align-items: center;
	/* Align items vertically centered */
	gap: 8px;
	/* Add some spacing between the elements */
}

.user {
	font-size: 14px;
	/* Adjust as needed */
}

.edit-button {
	margin-left: 8px;
	/* Optional if more spacing is needed */
}

.popup-overlay {
	position: fixed;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background: rgba(0, 0, 0, 0.5);
	display: flex;
	align-items: center;
	justify-content: center;
	z-index: 1000;
}


.emoji-buttons {
	display: flex;
	justify-content: center;
	gap: 10px;
	margin-bottom: 20px;
}

.emoji-button {
	background: #f0f0f0;
	border: none;
	border-radius: 50%;
	width: 50px;
	height: 50px;
	font-size: 24px;
	cursor: pointer;
	transition: transform 0.2s, background 0.2s;
	display: flex;
	align-items: center;
	justify-content: center;
}

.emoji-button:hover {
	background: #dfe6e9;
	transform: scale(1.2);
}

.popup-actions {
	margin-top: 10px;
}

.close-button {
	background: #e74c3c;
	color: #fff;
	border: none;
	padding: 10px 20px;
	border-radius: 5px;
	cursor: pointer;
	font-size: 16px;
}

.close-button:hover {
	background: #c0392b;
}
</style>
