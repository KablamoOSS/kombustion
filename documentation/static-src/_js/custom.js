(function() {
  $(document).ready(function(){
    if (localStorage.getItem("mode") == "dark") {
      $("body").addClass("dark");
    }
  });
  function toggleDarkMode() {
    if (localStorage.getItem("mode") !== "dark") {
      $("body").addClass("dark");
      localStorage.setItem("mode", "dark");
    } else {
      $("body").removeClass("dark");
      localStorage.setItem("mode", "light");
    }
  }
  $("#dark-mode-toggle").click(toggleDarkMode);
})();
