<template>
    <div>
        <v-btn
            elevation="2"
            @click="getTask"
            class="ma-5"
        >Get</v-btn>
        <v-btn
            elevation="2"
            @click="getTest"
            class="ma-5"
        >Test</v-btn>
        <v-btn
            elevation="2"
            @click="getSample"
            class="ma-5"
        >Sample</v-btn>
        <v-text-field
        label="New Task"
        hide-details="auto"
        v-model="task"
        ></v-text-field>
        <v-btn
        elevation="2"
        @click="addTask"
        class="ma-3"
        >Add</v-btn>
        <v-card
            class="mx-auto"
            max-width="300"
            tile
        >
            <v-list dense>
            <v-subheader>My Tasks</v-subheader>
            <v-list-item-group
                color="primary"
            >
                <v-list-item
                v-for="(item, i) in items"
                :key="i"
                >
                    <v-list-item-content>
                        <v-list-item-title v-text="item.text"></v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list-item-group>
            </v-list>
        </v-card>
    </div>
</template>

<script>
    import axios from 'axios'
    
    export default {
        data: () => ({
            items: [],
            tasks: [],
            task: ''
        }),
        methods:{
            addTask(){
                if(this.task.length != 0)this.items.push({"text":this.task});
                this.task = '';
            },
            getTask(){
                axios.get('/todo/get/')
                .then(response => {
                    if(response.status != 200){
                        console.log('error');
                        throw new Error('response error');
                    }else{
                        var results = response.data;
                        for(var i = 0;i < results.length;i++){
                            if(!results[i].Islimited){
                                this.items.push({"text":results[i].Title});
                                this.tasks.push(results[i])
                            }

                        }
                        console.log(this.tasks);
                    }
                })
                .catch(error =>{
                    alert(error)
                })
            },
            getTest(){
                axios.get('/test/')
                .then(response => {
                    console.log(response);
                })
            },
            getSample(){
                axios.get('https://jsonplaceholder.typicode.com/users')
                .then(response => {
                    var users = response.data
                    alert(users.name)
                    console.log(users)
                })
                .catch(error =>{
                    alert(error)
                })
            }
        }
    }
</script>