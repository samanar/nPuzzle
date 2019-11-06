<template>
  <v-app dark>
<!--    <v-navigation-drawer-->
<!--      v-model="drawer"-->
<!--      fixed-->
<!--      app-->
<!--      color="primary"-->
<!--      dark-->
<!--    >-->
<!--      <v-list>-->
<!--        <v-list-item-->
<!--          v-for="(item, i) in items"-->
<!--          :key="i"-->
<!--          :to="item.to"-->
<!--          router-->
<!--          exact-->
<!--        >-->
<!--          <v-list-item-action>-->
<!--            <v-icon>{{ item.icon }}</v-icon>-->
<!--          </v-list-item-action>-->
<!--          <v-list-item-content>-->
<!--            <v-list-item-title v-text="item.title"/>-->
<!--          </v-list-item-content>-->
<!--        </v-list-item>-->
<!--      </v-list>-->
<!--    </v-navigation-drawer>-->
<!--    <v-app-bar-->
<!--      color="primary"-->
<!--      dark-->
<!--      fixed-->
<!--      app-->
<!--    >-->
<!--      <v-app-bar-nav-icon @click.stop="drawer = !drawer"/>-->
<!--      <div class="logo_container">-->
<!--        <v-img src="/images/robot.svg" contain height="50" width="100px"></v-img>-->
<!--      </div>-->
<!--      <v-toolbar-title style="font-family: 'Roboto Light',serif"><h3>{{title}}</h3></v-toolbar-title>-->
<!--      <v-spacer/>-->
<!--    </v-app-bar>-->
    <v-content class="content_background">
      <v-container>
        <nuxt/>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
    export default {
        data() {
            return {
                drawer: false,
                title: 'Samanar',
                items: [
                    {
                        icon: 'mdi-apps',
                        title: 'Welcome',
                        to: '/'
                    },
                    {
                        icon: 'mdi-chart-bubble',
                        title: 'Inspire',
                        to: '/inspire'
                    }
                ],
            }
        }
    }
</script>

<style scoped>
  .content_background {
    background-color: #f0f0f0;
  }

  .logo_container {
    height: 100% !important;
    width: 70px !important;
  }
</style>
