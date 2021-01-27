<template>
  <v-card>
        <v-sheet tile height="6vh" color="grey lighten-3" class="d-flex align-center">
          <v-btn outlined small class="ma-4" @click="setToday">
            今日
          </v-btn>
          <v-btn icon @click="$refs.calendar.prev()">
            <v-icon>mdi-chevron-left</v-icon>
          </v-btn>
          <v-btn icon @click="$refs.calendar.next()">
            <v-icon>mdi-chevron-right</v-icon>
          </v-btn>
          <v-toolbar-title>{{ title }}</v-toolbar-title>
        </v-sheet>
        <v-sheet height="94vh">
          <v-calendar
            ref="calendar"
            v-model="value"
            :events="events"
            :event-color="getEventColor"
            locale="ja-jp"
            :day-format="(timestamp) => new Date(timestamp.date).getDate()"
            :month-format="(timestamp) => new Date(timestamp.date).getMonth() + 1 + ' /'"
            @change="getEvents"
            @click:event="showEvent"
            @click:date="viewDay"
          ></v-calendar>
        </v-sheet>

  </v-card>
</template>

<script>
import moment from 'moment';

export default {
  data: () => ({
    events: [],
    value: moment().format('yyyy-MM-DD'),
  }),
  computed: {
    title() {
      return moment(this.value).format('yyyy年 M月');
    },
  },
  methods: {
    setToday() {
      this.value = moment().format('yyyy-MM-DD');
    },
    getEvents() {
      const events = [
        // new Dateからmomentに変更
      ];
      this.events = events;
    },
    getEventColor(event) {
      return event.color;
    },
  },
};
</script>