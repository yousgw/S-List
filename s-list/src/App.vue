<template>
  <v-app>
    <v-container>
      <NewTask/>
      <v-row class="grey lighten-3">
        <v-col cols="3" style="background-color: #FFCDD2">
         <Test/>
        </v-col>
        <v-col cols="9" style="background-color: #F8BBD0">
          <HelloWorld/>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import HelloWorld from './components/HelloWorld';
import Test from './components/Test';
import NewTask from './components/NewTask';

export default {
  name: 'App',
  components: {
    HelloWorld,
    Test,
    NewTask,
  },

  data: () => ({
    //
  }),
};
</script>
