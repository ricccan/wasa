<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			token: localStorage.getItem("token"),
			username: localStorage.getItem("username"),
			chats: {},
			id: localStorage.getItem("id"),
			dynamicData: {},
			showPopup: false,
			newGroupname: null
		}
	},
	methods: {
		async newPage() {
			this.loading = true;
			this.errormsg = null;
			try {
				this.$router.push({ path: "/profile" });
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async login() {
			this.loading = true;
			this.errormsg = null;
			try {
				this.$router.push({ path: "/" });
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getChat() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/" + this.id + "/conversations", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token"),
					}

				}); // crea un json che gli passa un nome
				this.chats = response.data;
				this.createDynamicListsFromJSON(this.chats);
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		createDynamicListsFromJSON(json) {
			// Dynamically assign the given JSON to the dynamicData property
			this.dynamicData = json;
			if (this.dynamicData) {
				this.dynamicData.forEach((item, index) => {
					item.GroupPhoto = "data:image/jpeg;base64," + item.GroupPhoto;
				});
			}

		},
		closePopup() {
			this.showPopup = false;
		},

		async creaGruppo(){
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
		}

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
							<button type="button" class="btn btn-sm btn-outline-primary" @click="login">
								new chat
							</button>
						</div>
					</div>

					<!-- Scrollable List -->
					<div class="list-container">
						<ul>
							<li v-for="(item, index) in chats" :key="index" @click="handleClick(item)">
								<a class="clickable-item">
									<img :src="item.GroupPhoto" alt="Missing Photo" class="group-photo">
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
</style>
