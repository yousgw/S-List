<template>
    <dev> 
        <v-row>
            <v-col>
                <v-text-field
                label="New Task"
                hide-details="auto"
                v-model="task"
                class="ma-5"
                ></v-text-field>
            </v-col>
            <v-col cols="2">
                <v-btn
                elevation="2"
                @click="addTask"
                class="ma-5"
                >Add</v-btn>
            </v-col>
        </v-row>
        <v-card>
            <v-expansion-panels>
                <v-expansion-panel                >
                <v-expansion-panel-header>
                    詳細設定
                </v-expansion-panel-header>
                <v-expansion-panel-content>
                    <v-switch
                    class="ma-3"
                    v-model="limit"
                    label="期限指定"
                    color="red"
                    @click="show=!show"
                    ></v-switch>
                    <div
                        v-show="show"
                    >
                        <v-menu
                            ref="menu"
                            v-model="menu"
                            :close-on-content-click="false"
                            :return-value.sync="date"
                            transition="scale-transition"
                            offset-y
                            min-width="290px"
                        >
                            <template v-slot:activator="{ on, attrs }">
                            <v-text-field
                                v-model="date"
                                label="日付を指定"
                                prepend-icon="mdi-calendar"
                                readonly
                                v-bind="attrs"
                                v-on="on"
                            ></v-text-field>
                            </template>
                            <v-date-picker
                            v-model="date"
                            no-title
                            scrollable
                            >
                            <v-spacer></v-spacer>
                            <v-btn
                                text
                                color="primary"
                                @click="menu = false"
                            >
                                Cancel
                            </v-btn>
                            <v-btn
                                text
                                color="primary"
                                @click="$refs.menu.save(date)"
                            >
                                OK
                            </v-btn>
                            </v-date-picker>
                        </v-menu>
                    </div>
                    <v-switch
                    class="ma-3"
                    v-model="share"
                    label="公開設定"
                    ></v-switch>
                </v-expansion-panel-content>
                </v-expansion-panel>
            </v-expansion-panels>
        </v-card>
        
        <v-card
            class="ma-5"
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
    </dev>
</template>

<script>
    export default {
        data: () => ({
            items: [],
            task: '',
            date: new Date().toISOString().substr(0, 10),
            menu: false,
            show: false,
        }),
        methods:{
            addTask(){
                if(this.task.length != 0){
                    var info = '';
                    this.items.push({"text":this.task+this.date+this.share});
                    info = "新規タスク"+this.task;

                    if (this.show) info += "\n期限："+this.date;
                    else info += "\n期限なし";
                    if (this.share) info += "\n公開";
                    else info += "\n非公開";

                    alert(info);
                
                    this.task = '';
                }
                
            },
        }
    }
</script>