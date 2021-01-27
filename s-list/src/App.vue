<template>
  <v-app>
    <v-container>
      <NewTask/>
      <v-row class="grey lighten-3">
        <v-col cols="3" style="background-color: #FFCDD2">
         <List/>
        </v-col>
        <v-col cols="9" style="background-color: #F8BBD0">
          <Calendar/>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import Calendar from './components/Calendar';
import List from './components/List';
import NewTask from './components/NewTask';

export default {
  name: 'App',
  components: {
    Calendar,
    List,
    NewTask,
  },

  data: () => ({
    //
  }),
};
</script>
