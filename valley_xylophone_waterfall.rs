fn main(){
    //Prints a welcome message about Green Thumbs
    println!("Welcome to Green Thumbs! A place to learn about gardening and sustainability!");

    //Declares variabels
    let seed_count: u32 = 10;
    let soil_type: &str = "loam";
    let seed_type: &str = "sunflower";
    let water_level: f64 = 0.75;

    //Sets up a loop to ask the user for input on where to start
    println!("Where would you like to start?");
    println!("1. Shopping for seeds");
    println!("2. Preparing the soil");
    println!("3. Planting the seeds");
    println!("4. Watering the plants");
    println!("5. Harvesting the plants");

    let mut user_selection = String::new();
    io::stdin().read_line(&mut user_selection)
        .expect("Failed to read line");

    //Shopping for Seeds
    if user_selection.trim() == "1" {
        println!("Let's go shopping for seeds! You need {} {} seeds.", seed_count, seed_type);
        println!("You can find the seeds at your local nursery or garden center.");
    }
    //Preparing the Soil
    else if user_selection.trim() == "2" {
        println!("Let's get the soil ready for planting. You will need {} soil.", soil_type);
        println!("You can find {} soil at your local nursery or garden center.", soil_type);
    }
    //Planting the Seeds
    else if user_selection.trim() == "3" {
        println!("Let's plant the seeds! Take {} {} seeds and place them in the soil.", seed_count, seed_type);
        println!("Make sure to space them evenly so that they have enough room to grow.");
    }
    //Watering the Plants
    else if user_selection.trim() == "4" {
        println!("Let's water the plants. The soil should be kept at a moisture level of {}.", water_level);
        println!("You can use a watering can or garden hose to water the plants.");
    }
    //Harvesting the Plants
    else if user_selection.trim() == "5" {
        println!("Let's harvest the plants! When the {} seeds are ready, you can harvest them.", seed_type);
        println!("You can use scissors or your hands to gently remove the seeds.");
    }
    //Invalid Selection
    else {
        println!("Invalid selection. Please choose a number between 1 and 5.");
    }
}