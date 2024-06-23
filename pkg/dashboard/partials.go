package dashboard

import views_variables "github.com/ppp3ppj/pppfav-htmx/views/variables"

func nav(indexSelected int) []views_variables.DrawerMenu {
    lists := []views_variables.DrawerMenu{
        {
            Label: "Persons",
            URL: "/dashboard/persons",
            IsActive: false,
            IsBoosted: true,
        },
        {
            Label: "Back to Site",
            URL: "/",
            IsActive: false,
            IsBoosted: true,
        },
    }

    lists[indexSelected].IsActive = true

    return lists
}
