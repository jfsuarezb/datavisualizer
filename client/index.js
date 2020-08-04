Vue.component('user-panel-display', {

    data:function() {

        return {
            uid: app.$route.params.uid
        }

    },

    template: `
    
        <h1>{{ this.uid }}</h1>

    `

})

Vue.component('date-picker', {

    data: function() {
        return {
            date: "",
            buttonState:false,
            usersData: []
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

    router,
    vuetify: new Vuetify(),

}).$mount("#app")