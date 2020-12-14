<template>
    <div>
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
    export default {
        data: () => ({
            items: [],
            task: ''
        }),
        methods:{
            addTask(){
                if(this.task.length != 0)this.items.push({"text":this.task});
                this.task = '';
            }
        }
    }
</script>