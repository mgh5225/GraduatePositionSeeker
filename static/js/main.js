window.addEventListener("load", () => {
  const univ_select = document.getElementById("universityID");
  const fac_select = document.getElementById("facultyID");

  if (univ_select != undefined && fac_select != undefined)
    univ_select.onchange = async (e) => {
      const u_id = e.target.value;

      const data = await fetch(`/api/universities/${u_id}/faculties`);

      const faculties = await data.json();

      while (fac_select.options.length > 0) {
        fac_select.remove(0);
      }

      for (const faculty of faculties) {
        const newOption = new Option(faculty.name, faculty.id);
        fac_select.add(newOption, undefined);
      }
    };
});
