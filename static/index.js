fetch("/banners").then((resp) => {
  resp.json().then((banners) => {
    const div = document.querySelector(".section div");
    for (const banner of banners) {
      div.innerHTML += `<label class="radio">
            <input type="radio" name="banner" class="radio" value="${banner}">
            <h2>${banner}</h2>
         </label>`;
    }
  });
});
