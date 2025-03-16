use bevy::prelude::*;

// Components
#[derive(Component)]
struct Player {
    speed: f32,
}

#[derive(Component)]
struct Camera;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins.set(WindowPlugin {
            primary_window: Some(Window {
                title: "Marisalt: A Dark Fantasy Pirate Role Playing Game".to_string(),
                ..default()
            }),
            ..default()
        }))
        .add_systems(Startup, setup)
        .add_systems(Update, (camera_follow, player_movement))
        .run();
}

fn setup(mut commands: Commands) {
    // Spawn a 2D camera
    commands.spawn((Camera2d::default(), Camera));

    // Spawn the player
    commands.spawn((
        Sprite {
            color: Default::default(),
            custom_size: Some(Vec2::new(50.0, 50.0)),
            ..default()
        },
        Transform {
            translation: Vec3::new(0.0, 0.0, 1.0),
            scale: Vec3::new(1.0, 1.0, 1.0),
            rotation: Quat::IDENTITY,
        },
        Player { speed: 150.0 },
    ));

    // Spawn a yellow building placeholder
    commands.spawn((
        Sprite {
            color: Color::linear_rgba(1.0, 0.8, 0.2, 0.5),
            custom_size: Some(Vec2::new(100.0, 100.0)),
            ..default()
        },
        Transform {
            translation: Vec3::new(200.0, 200.0, 0.0),
            ..default()
        },
    ));
}

fn camera_follow(
    player_query: Query<&Transform, With<Player>>,
    mut camera_query: Query<&mut Transform, (With<Camera>, Without<Player>)>,
    time: Res<Time>,
) {
    if let Ok(player_transform) = player_query.get_single() {
        if let Ok(mut camera_transform) = camera_query.get_single_mut() {
            let lerp_speed = 5.0;
            camera_transform.translation = camera_transform
                .translation
                .lerp(player_transform.translation, lerp_speed * time.delta_secs());
        }
    }
}

fn player_movement(
    keyboard_input: Res<ButtonInput<KeyCode>>,
    time: Res<Time>,
    mut player_query: Query<(&mut Transform, &Player)>,
) {
    if let Ok((mut transform, player)) = player_query.get_single_mut() {
        let mut direction = Vec2::ZERO;

        // Handle WASD for movement
        if keyboard_input.pressed(KeyCode::KeyW) {
            direction.y += 1.0;
        }
        if keyboard_input.pressed(KeyCode::KeyS) {
            direction.y -= 1.0;
        }
        if keyboard_input.pressed(KeyCode::KeyA) {
            direction.x -= 1.0;
        }
        if keyboard_input.pressed(KeyCode::KeyD) {
            direction.x += 1.0;
        }

        // Normalize the direction vector and apply movement
        if direction.length_squared() > 0.0 {
            direction = direction.normalize();
            transform.translation.x += direction.x * player.speed * time.delta_secs();
            transform.translation.y += direction.y * player.speed * time.delta_secs();
        }
    }
}
