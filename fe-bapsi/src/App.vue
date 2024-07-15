<template>
  <NavBar v-if="isAuth" />
  <!-- <NavBarAdmin v-if="isAuth && role == 'admin'" /> -->

  <div class="container">
    <router-view />
  </div>
  <FooterBar />
</template>

<script>
import NavBar from "@/components/navigation/NavBar.vue";
// import NavBarAdmin from "@/components/navigation/NavBar-Admin.vue";
import FooterBar from "@/components/navigation/FooterBar.vue";
// import {axios} from "@/axios/config.js";

export default {
  components: {
    NavBar,
    // NavBarAdmin,
    FooterBar,
  },

  data() {
    return {
      role: "",
    }
  },

  computed: {
    isAuth() {
      // Check if the current route is one of the specified pages
      const validPages = [
        'HomeView',
        'ProfileView',
        'ProfileOtherView',
        'InvitationView',
        'KonfigurasiView',
        'DaftarStaff',
        'TambahLoket',
        'TambahDivisi',
        'AbsenDaruratView',
      ];
      return validPages.includes(this.$route.name);
    },
  },

  methods: {
    PushRoute(pageName) {
      // Smooth scroll to top of the page
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });

      this.$router.push({ name: pageName });
    }
  }
}
</script>

<style>
@import "@/stylesheets/typography.css";
@import "@/stylesheets/coloring.css";
@import "@/stylesheets/form.css";
@import "@/stylesheets/table.css";
@import "@/stylesheets/positioning.css";
@import "@/stylesheets/button/bt-1.css";
@import "@/stylesheets/button/bt-2.css";
@import "@/stylesheets/button/bt-3.css";

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}

/* WebKit */
::-webkit-scrollbar {
  display: none;
  width: 10px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
}

::-webkit-scrollbar-thumb {
  background: #888;
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
}
</style>
