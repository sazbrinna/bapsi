<template>
  <div class="card-pesan">
    <p>Ini adalah pesan untuk diketahui bersama</p>
    <div class="profile row">
      <img src="@/assets/imgs/profile.jpg" alt="profile">
      <div class="profile-data">
        <h5>Faizal Maulana</h5>
        <p class="quote">System Development Center</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* style for card pesan */
.card-pesan {
  gap: 1.5rem;
  box-shadow: 0 0 7px 0 #888888;
  padding: 1.5rem 2rem;
  border-radius: 8px;
  align-items: flex-start;
  width: 25%;
}

.card-pesan .profile-data {
  align-items: flex-start;
}

.card-pesan h5,
.card-pesan p.quote {
  text-align: start;
}

.card-pesan img {
  width: 45px;
  height: 45px;
  object-fit: cover;
  border-radius: 5rem;
}

.card-pesan .profile {
  gap: 1rem;
}
</style>