package postgres

const CreateUsersTable  =  `Create table if not exists users (
	id bigserial primary key,
	name varchar(60) not null,
	surname varchar(60) not null,
	login varchar(60) not null unique,
	password  text ,
	email varchar(60) unique,
	phone varchar(60) unique,
	status boolean not null,
	created_at timestamp default CURRENT_TIMESTAMP
);`

const CreateCitiesTable  =  `Create table if not exists cities (
	id bigserial primary key,
	name varchar(60) not null unique
);`

const CreateBranchesTable  =  `Create table if not exists branches (
	id bigserial primary key,
	address varchar(60) not null unique,
	city_id integer references cities
);`

const CreateRolesTable  =  `Create table if not exists roles (
	id bigserial primary key,
	name varchar(60) not null unique,
	display_name varchar(60) not null,
	description varchar(191) not null
)`

const CreateUserRoleTable  =  `Create table if not exists userRole (
	role_id integer references roles,
	user_id integer references users unique
);`

const CreatePurposesTable  =  `Create table if not exists purposes (
	id bigserial primary key,
	name varchar(90) not null unique
);`

const CreateTimesTable  =  `Create table if not exists times (
	id bigserial primary key,
	name varchar(30) not null unique
);`

const CreateQueuesTable  =  `Create table if not exists queues (
	id bigserial primary key,
	queue_code integer not null,
	terminal_id integer references terminals("terminal_number"),
	user_id integer references users,
	city_id integer references cities,
	branch_id integer references branches,
	purpose_id integer references purposes,
	time_id integer references times,
	status varchar(60) not null,
	date date default current_date,
	start_at timestamp,
	finish_at timestamp,
	created_at timestamp default CURRENT_TIMESTAMP
);`

const CreateTerminalsTable  =  `Create table if not exists terminals (
	id bigserial primary key,
	terminal_number integer not null unique,
	city_id integer references cities,
	branch_id integer references branches,
	user_id integer references users
);`
