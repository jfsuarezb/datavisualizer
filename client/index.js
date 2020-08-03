Vue.component('date-picker', {

    data: function() {
        return {
            date: ""
        }
    },

    template:`
    
        <div id=\"datepicker\">
            <h1 id=\"datepickerbanner\">Selecciona una Fecha: </h1>
            <input type="date" v-model="date"></input>
            <v-btn small id="datepickerbutton" v-on:click="clickHandler">Buscar</v-btn>
        </div>

    `,

    methods:{

        clickHandler: async function() {
            let unixt = Date.parse(this.date)/1000
            try {
                await fetch(`http://localhost:3000/sync?date=${unixt}`)
                let resp = await fetch(`http://localhost:3000/buyers`)
                app.usersData = resp.json()
                console.log(app.usersData)
            } catch(e) {
                alert(e)
            }
        }

    }

})

var app = new Vue({

    el: "#app",
    data: {
        usersData: {},
        userData:{}
    },
    vuetify: new Vuetify(),

})