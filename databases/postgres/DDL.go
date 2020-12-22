package postgres

const CreateUsersTable  =	`Create table if not exists users (
	id bigserial primary key,
	name varchar(60) not null,
	surname varchar(60) not null,
	login varchar(60) not null unique,
	password  text not null,
	email varchar(60) unique,
	phone varchar(60) unique,
	status boolean not null,
	created_at time default CURRENT_TIMESTAMP
);`

const CreateCitiesTable  =  `Create table if not exists cities (
	id bigserial primary key,
	name varchar(60)
);`

const CreateBranchesTable  =  `Create table if not exists branches (
	id bigserial primary key,
	address varchar(60) not null,
	city_id integer references cities
);`

const CreateRolesTable  =  `Create table if not exists roles (
	id bigserial primary key,
	name varchar(60) not null,
	display_name varchar(60) not null,
	description varchar(191) not null
`

const CreateUserRoleTable  =  `Create table if not exists userRole (
	role_id references roles,
	user_id references users
);`

const CreatePurposesTable  =  `Create table if not exists purposes (
	id bigserial primary key,
	name varchar(90) not null
);`

const CreateTimesTable  =  `Create table if not exists times (
	id bigserial primary key,
	name varchar(30) not null
);`

const CreateQueuesTable  =  `Create table if not exists queues (
	id bigserial primary key,
	user_id references users,
	purpose_id references purposes,
	time_id references times,
	city_id references cities,
	branch_id references branches,
	status string not null,
	date time,
	start_at time,
	finish_at time,
	created_at time default CURRENT_TIMESTAMP
);`
