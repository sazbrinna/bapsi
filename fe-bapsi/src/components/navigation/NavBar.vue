<template>
  <!-- Role Asisten -->

  <nav :class="{ hidden: isNavHidden }" v-if="role == 'asisten'">
    <div class="wrapper row">
      <div class="logo">
        <img src="@/assets/logo.svg" alt="Bapsi" @click="PushRoute('HomeView')">
      </div>
      <div class="navigation row">
        <div class="option row phone-hide">
          <h5 @click="PushRoute('HomeView')">Home</h5>
          <h5 @click="PushRoute('ProfileView')">Profile</h5>
          <h5 @click="logout()" style="color: var(--red);">Keluar</h5>
        </div>

        <button class="bt-1" @click="submitMasuk()" v-if="showAbsen && buttonAbsen == 'notCheck'">Absen Masuk</button>
        <button class="bt-2" @click="submitKeluar()" v-if="showAbsen && buttonAbsen == 'checkedIn'">Absen
          Keluar</button>


        <div class="hamburger" @click.prevent="handleHamburger()">
          <img src="@/assets/icons/menu.png" alt="menu">
        </div>
      </div>
    </div>
  </nav>

  <!-- Role Admin -->

  <nav v-if="role == 'admin'">
    <div class="wrapper row">
      <div class="logo">
        <img src="@/assets/logo.svg" alt="Bapsi" @click="PushRoute('HomeView')">
      </div>
      <div class="navigation row">
        <div class="option row phone-hide">
          <h5 @click="PushRoute('HomeView')">Home</h5>
          <h5 @click="PushRoute('DaftarStaff')">Asisten</h5>
          <h5 @click="PushRoute('KonfigurasiView')">Konfigurasi</h5>
          <h5 @click="logout()" style="color: var(--red);">Keluar</h5>
          <h5 @click="PushRoute('AbsenDaruratView')">Absen Darurat</h5>
        </div>

        <div class="hamburger" @click.prevent="handleHamburger()">
          <img src="@/assets/icons/menu.png" alt="menu">
        </div>
      </div>
    </div>
  </nav>

  <!-- navigation on phone (fullscreen) -->
  <nav :class="{ openHamburger: isHamburgerShowing, openHamburgerHidden: initialHamburger }">
    <div class="wrapper">
      <div class="logo">
        <img src="@/assets/logo.svg" alt="Bapsi" @click="PushRoute('HomeView')">
      </div>
      <div class="navigation column">
        <div class="option column">
          <h5 @click="PushRoute('HomeView')">Home</h5>
          <h5 v-if="role != 'admin'" @click="PushRoute('ProfileView')">Profile</h5>

          <h5 v-if="role == 'admin'" @click="PushRoute('DaftarStaff')">Asisten</h5>
          <h5 v-if="role == 'admin'" @click="PushRoute('KonfigurasiView')">Konfigurasi</h5>
          <h5 v-if="role == 'admin'" @click="PushRoute('AbsenDaruratView')">AbsenDarurat</h5>

          <h5 @click="handleHamburger()">Back</h5>
        </div>

      </div>
    </div>
  </nav>

  <AlertCard :msg="message" v-if="message" />
</template>

<script>
import { axios } from "@/axios/config.js";
import AlertCard from "@/components/alert/AlertCard.vue";

export default {
  data() {
    return {
      latitude: "",
      longitude: "",
      status: false,
      lokasi: "",
      showAbsen: true,
      buttonAbsen: "",
      email: "",

      role: "",

      // NavBehavior
      isNavHidden: false,
      lastScrollPosition: 0,
      showNavMenu: false,
      isHamburgerShowing: false,
      initialHamburger: true,

      message: "",
    }
  },

  components: {
    AlertCard,
  },

  beforeCreate() {
    axios.get("/absen/cek")
      .then((response) => {
        let checkinStatus = response.data.message;
        switch (checkinStatus) {
          case "Checked-In":
            this.buttonAbsen = "checkedIn";
            break;
          case "Checked-Out":
            this.buttonAbsen = "checkedOut";
            break;
          default:
            this.buttonAbsen = "notCheck";
            break;
        }
      })
  },

  created() {
    this.getLocation();
    this.getRole();
    window.addEventListener("scroll", this.handleScroll);
  },

  beforeMount() {
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    });
  },

  mounted() {
    // this.otomaticMasuk();

    this.intervalId = setInterval(() => {
      if (this.buttonAbsen == "notCheck") {
        this.getLocation();
        // this.otomaticMasuk();
      } else {
      // Stop the interval when condition is false
      clearInterval(this.intervalId);
      }
    }, 5000);
  },

  beforeUnmount() {
    // Clear the interval when the component is destroyed to prevent memory leaks
    clearInterval(this.intervalId);
  },

  methods: {
    handleHamburger() {
      this.isHamburgerShowing = !this.isHamburgerShowing;
      this.initialHamburger = !this.initialHamburger;
    },

    logout() {
      localStorage.clear();
      window.location.reload();
    },

    getLocation() {
      if ("geolocation" in navigator) {
        navigator.geolocation.getCurrentPosition(
          (position) => {
            // Extract latitude and longitude from the position object
            const latitude = position.coords.latitude;
            const longitude = position.coords.longitude;

            // // Get current time
            // const currentTime = new Date();
            // // Get current hour
            // const currentHour = currentTime.getHours();

            // // Check if it's before 16:00 (4:00 PM)
            // const isBefore16 = currentHour < 16;

            // Update component data based on conditions
            this.latitude = latitude;
            this.longitude = longitude;
            // this.showAbsen = isBefore16; // Update showAbsen based on current time

            // Optionally, you may emit an event or call another method here
            // to notify parent components or perform additional actions.

          },
          (error) => {
            console.error("Error getting geolocation:", error);
            // Optionally, display an error message to the user or handle it gracefully.
          }
        );
      } else {
        console.error("Geolocation is not supported by this browser.");
        // Optionally, display an error message to the user or handle it gracefully.
      }
    },


    getRole() {
      // handler for get the role of the user for deciding which nav will be used for

      axios.get("/role")
        .then((response) => {
          console.log(response);
          this.role = response.data.role;
        })
    },

    handleScroll() {
      // handler for navigation scroll

      const currentScrollPosition = window.pageYOffset;

      if (currentScrollPosition > this.lastScrollPosition) {
        this.isNavHidden = true;
      } else {
        this.isNavHidden = false;
      }

      this.lastScrollPosition = currentScrollPosition
    },

    PushRoute(pageName) {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });

      this.handleHamburger();

      this.$router.push({ name: pageName });
    },

    submitKeluar() {
      // handler for absen out

      const data = {
        latitude: this.latitude.toString(),
        longitude: this.longitude.toString(),
      };

      axios.post("/absen/keluar", data)
        .then((response) => {
          const lokasi = response.data.lokasi;
          this.lokasi = lokasi;
          this.status = true;
          window.location.reload();
        })
        .catch((err) => {
          const errorCode = err.response.status
          if (errorCode == 406) {
            this.message = "Kamu berada di luar area loket";
            setTimeout(() => {
              this.message = "";
            }, 2000)
          }
        })
    },

    submitMasuk() {
      // handler for absen in
      const data = {
        latitude: this.latitude.toString(),
        longitude: this.longitude.toString(),
        lokasi: this.lokasi,
      };

      axios.post("/absen/masuk", data)
        .then(() => {
          window.location.reload();
        })
        .catch((err) => {
          const errorCode = err.response.status
          if (errorCode == "406") {
            this.message = "Kamu berada di luar area loket";
            setTimeout(() => {
              this.message = "";
            }, 2000)
          }
        })
    },

    // otomaticMasuk() {
    //   const data = {
    //     latitude: this.latitude.toString(),
    //     longitude: this.longitude.toString(),
    //     lokasi: this.lokasi,
    //   };

    //   axios.post("/absen/masuk", data)
    //     .then(() => {
    //       window.location.reload();
    //     })
    //     .catch(() => {
    //     })
    // }
  }
}
</script>

<style scoped>
nav {
  height: 62px;
  background-color: white;
  overflow: hidden;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  z-index: 9999;
  overflow: hidden;
  transition: 0.5s;
}

nav.hidden {
  transform: translateY(-100%);
}

nav.openHamburger {
  display: none;
}

.openHamburgerHidden {
  display: none !important;
}

.wrapper {
  align-items: center;
  justify-content: space-between;
}

.navigation {
  gap: 1rem;
  align-items: center;
}

.hamburger {
  display: none;
}

button {
  min-width: 136px;
}

h5 {
  width: 100%;
  text-align: center;
  text-wrap: nowrap;
  cursor: pointer;
}

h5:hover {
  color: var(--primary);
}

.logo {
  display: flex;
  align-items: center;
  justify-content: center;
}

img {
  height: 32px;
}

@media screen and (max-width: 768px) {
  nav .wrapper {
    justify-content: space-between;
    align-items: center;
    flex-wrap: nowrap !important;
  }

  .navigation {
    justify-content: center;
    gap: 1em;
  }

  .navigation .phone-hide {
    display: none;
  }

  .navigation .hamburger {
    display: block;
  }

  nav.openHamburgerHidden {
    display: none;
  }

  nav.openHamburger {
    position: fixed;
    /* Use fixed positioning to keep it fixed relative to the viewport */
    top: 50%;
    /* Position at the middle of the viewport */
    left: 0;
    /* Start from the left edge of the viewport */
    transform: translateY(-50%);
    /* Adjust to center vertically */
    display: flex !important;
    flex-direction: column;
    /* Ensure items stack vertically */
    width: 100%;
    height: 100vh;
    z-index: 9998;
    /* Ensure it appears above other content */
  }

  nav.openHamburger .logo {
    padding-bottom: 1em;
    border-bottom: 1px solid var(--secondary)
  }

  nav.openHamburger .option {
    gap: 1rem;
  }

  .hamburger {
    display: block;
  }
}
</style>