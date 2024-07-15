<template>
  <footer class="gradient">
    <div class="wrapper">
      <div class="content-footer row">
        <div class="part-content">
          <h2 @click.prevent="PushRoute('HomeView')">BAPSI GUNADARMA</h2>
        </div>
        <div class="part-content gap">
          <div class="navigation">
            <h3>Quick Links</h3>
            <a href="https://gunadarma.ac.id/" target="_blank">Gunadarma</a>
            <a href="https://studentsite.gunadarma.ac.id/" target="_blank">Studentsite</a>
            <a href="https://v-class.gunadarma.ac.id/" target="_blank">Virtual Class</a>
            <a href="http://www.career.gunadarma.ac.id/" target="_blank">UG Career Center</a>
            <a href="https://www.baak.gunadarma.ac.id/" target="_blank">BAAK</a>
            <a href="https://psa.gunadarma.ac.id/" target="_blank">PSA</a>
            <a href="https://staffsite.gunadarma.ac.id/" target="_blank">Staffsite</a>
          </div>
          <div class="navigation">
            <h3>Contact Us</h3>
            <span>
              <img src="@/assets/icons/place-mark.svg" alt="">
              <a href="https://maps.app.goo.gl/L5TCxqqS53dRSajm6" target="_blank">Jl. Margonda Raya - 100</a>
            </span>
            <span>
              <img src="@/assets/icons/email.svg" alt="">
              <a href="mailto:mediacenter.gunadarma.ac.id">mediacenter.gunadarma.ac.id</a>
            </span>
            <span>
              <img src="@/assets/icons/phone.svg" alt="">
              <a href="#">+62 (21) 78881112</a>
            </span>
          </div>
        </div>
        <div class="part-content row sosmed">
          <img src="@/assets/icons/socialmedia/InstagramLogo.svg" alt="">
          <img src="@/assets/icons/socialmedia/LinkedinLogo.svg" alt="">
          <img src="@/assets/icons/socialmedia/TiktokLogo.svg" alt="">
          <img src="@/assets/icons/socialmedia/TwitterLogo.svg" alt="">
        </div>
      </div>
      <div class="copyright">
        <p>@ copyright 2024, all right reserved</p>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  methods: {
    PushRoute(pageName) {
      // Smooth scroll to top of the page
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });

      this.$router.push({ name: pageName });
    },

    logout() {
      localStorage.clear();
    }
  }
}
</script>

<style scoped>
footer {
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

a {
  color: white;
}

.content-footer {
  display: flex;
  gap: 3rem;
  align-items: center;
  justify-content: space-around;
  padding: 2rem 0 2rem 0;
  flex-wrap: wrap;
}

.part-content {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.sosmed {
  gap: 0.6rem;
}

.navigation {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 0.5rem;
  flex-wrap: wrap;
}

span {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 7px;
}

.gap {
  gap: 2rem;
}

.copyright {
  border-top: 1px solid white;
  max-width: 700px;
  width: 100%;
}

/* image */
.sosmed img {
  width: 40px;
}

footer span img {
  width: 14px;
}

p {
  margin: 0.7rem 0 0.7rem 0;
}

@media screen and (max-width: 768px) {
  .navigation {
    align-items: center;
  }
}
</style>