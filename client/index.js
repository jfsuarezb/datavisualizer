Vue.component('user-panel-display', {

    data: function() {

        return {
            userName:"Julian",
            userAge:15,
            userTransHistory:[],
            otherUsers:[],
            recommendedProducts:[],
            buttonState:false
        }

    },

    template: `

        <div id="userPanelContainer">
            <v-btn style="font-size:1vw;" v-bind:disabled="buttonState" v-on:click="goBack" text small id="backButton">
                <v-icon small left>mdi-arrow-left</v-icon>Atras
            </v-btn>
            <h1 style="font-size:3vw;">Hola {{ this.userName }}!</h1>
            <div id="primaryInfoContainer">
                <p>Nombre: {{ this.userName }}</p>
                <p>Edad: {{ this.userAge }}</p>
            </div>
            <v-container>
                <v-row style="margin-top:1.5vw;">
                    <v-col style="padding:1vw;">
                        <v-card>
                            <v-toolbar style="background-color:rgb(55, 173, 112);color:white;">
                                <v-toolbar-title>Tus Transacciones</v-toolbar-title>
                            </v-toolbar>
                            <v-list>
                                <v-list-group v-for="trans in userTransHistory" :key="trans.ip">
                                    <template v-slot:activator>
                                        <v-list-item-content>
                                            <v-list-item-title v-text="'Transacción: ' + trans.ip"></v-list-item-title>
                                        </v-list-item-content>
                                    </template>
                                    <v-list-item v-for="prod in trans.prods" :key="prod.price">
                                        <v-list-item-content>
                                            <v-list-item-title v-text="'Producto: ' + prod.name"></v-list-item-title>
                                            <v-list-item-subtitle v-text="'Precio: ' + prod.price"></v-list-item-subtitle>
                                        </v-list-item-content>
                                    </v-list-item>
                                </v-list-group>
                            </v-list>
                        </v-card>
                    </v-col>
                    <v-col style="padding:1vw;">
                        <v-card style="margin-bottom:1vw;">
                            <v-toolbar style="background-color:rgb(55, 173, 112);color:white;">
                                <v-toolbar-title>Produtos Recomendados</v-toolbar-title>
                            </v-toolbar>
                            <v-list>
                                <v-list-item v-for="prod in recommendedProducts" :key="prod.price">
                                    <v-list-item-content>
                                            <v-list-item-title v-text="'Producto: ' + prod.name"></v-list-item-title>
                                            <v-list-item-subtitle v-text="'Precio: ' + prod.price"></v-list-item-subtitle>
                                    </v-list-item-content>
                                </v-list-item>
                            </v-list>
                        </v-card>
                        <v-card>
                            <v-toolbar style="background-color:rgb(55, 173, 112);color:white;">
                                <v-toolbar-title>Otros como tú</v-toolbar-title>
                            </v-toolbar>
                            <v-list>
                                <v-list-item v-for="user in otherUsers" :key="user.age">
                                    <v-list-item-content>
                                            <v-list-item-title v-text="'Nombre: ' + user.name"></v-list-item-title>
                                            <v-list-item-subtitle v-text="'Edad: ' + user.age"></v-list-item-subtitle>
                                    </v-list-item-content>
                                </v-list-item>
                            </v-list>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
        </div>

    `,
    methods: {

        getUserInfo: async function() {

            try {
                let resp = await fetch(window.location.href.replace(`#/user/${app.$route.params.uid}`,"") + `buyer?uid=${app.$route.params.uid}`)
                let data = await resp.json()
                return data
            } catch (error) {
                alert(error)
            }

        },
        
        goBack: function() {

            this.buttonState = true
            this.$router.push("/")

        }

    },
    created() {

        this.getUserInfo().then((data) => {
            console.log("hello")
            this.userName = data["getBuyer"][0]["name"]
            this.userAge = data["getBuyer"][0]["age"]
            this.userTransHistory = data["getBuyer"][0]["trans"]
            this.otherUsers = data["otherBuyersSameIp"]
            this.recommendedProducts = data["recommendedProducts"]
        }).catch((error) => {

            alert(error)

        })

    }

})

Vue.component('date-picker', {

    data: () => {
        
        if (typeof app === "undefined") {

            return {

                usersData:[],
                date:"",
                buttonState:false

            }

        } else {

            return {

                usersData: app.$data.usersData,
                date:"",
                buttonState:false

            }

        }

    },

    template:`
    
        <div id=\"datepicker\">
            <h1 id=\"datepickerbanner\">Selecciona una Fecha: </h1>
            <input type="date" v-model="date"></input>
            <v-btn v-bind:disabled="buttonState" small id="datepickerbutton" v-on:click="clickHandler">Buscar</v-btn>
            <v-list two-line subheader id="userlist">
                <v-list-item class="listClass" v-for="user in usersData" v-bind:key="user.uid">
                    <v-list-item-content>
                        <v-list-item-title v-text="'Nombre: ' + user.name"></v-list-item-title>
                        <v-list-item-subtitle v-text="'Edad: ' + user.age"></v-list-item-subtitle>
                    </v-list-item-content>
                    <v-list-item-action>
                        <v-btn v-on:click="userClickHandler(user.uid)" text small id="viewButton">Ver</v-btn>
                    </v-list-item-action>
                </v-list-item>
            </v-list>
        </div>

    `,

    methods:{

        clickHandler: function() {

            this.getBuyers().then(function(data) {
                app.$children[1].usersData = data["getBuyers"]
            }).catch(function(error) {
                alert(error)
            })

        },

        getBuyers: async function() {
            let unixt = Date.parse(this.date)/1000
            try {
                this.buttonState = true
                await fetch(window.location.href.replace("#/","") + `sync?date=${unixt}`)
                let resp = await fetch(window.location.href.replace("#/","") + "buyers")
                let data = await resp.json()
                this.buttonState = false
                return data
            } catch (error) {
                alert(error)
                this.buttonState = false
            }
        },

        userClickHandler: function(uid) {

            app.$data.usersData = app.$children[1].usersData
            this.$router.push(`user/${uid}`)

        }

    }

})

const routes = [
    {path:"/", component:'date-picker'},
    {path:"/user/:uid", component:'user-panel-display'}
]

Vue.use(VueRouter)

const router = new VueRouter({

    routes

})

var app = new Vue({

    data:{

        usersData:[]

    },
    router,
    vuetify: new Vuetify(),

}).$mount("#app")