CREATE TABLE IF NOT EXISTS person (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50) NULL,
    gender VARCHAR(6) NULL,
    age INT NULL,
    nationality VARCHAR(50) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ
);

CREATE OR REPLACE FUNCTION paginate(page_id int, 
                                         per_page int, 
                                         flt_name varchar(50), 
                                         flt_surname varchar(50),
                                         flt_patronymic varchar(50),
                                         flt_gender varchar(6),
                                         flt_age int,
                                         flt_nationality varchar(50)
                                        )
RETURNS TABLE (
  	id int,
  	name VARCHAR(50),
    surname VARCHAR(50),
    patronymic VARCHAR(50),
    gender VARCHAR(6),
    age INT,
    nationality VARCHAR(50),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ) AS $$
BEGIN
	if page_id < 1 then
  	page_id = 1;
  end if;

  RETURN query(
	select p.id, p.name, p.surname, p.patronymic, p.gender, p.age, p.nationality, p.created_at, p.updated_at from person p
  where (p.name = flt_name or flt_name = '') 
    and (p.surname = flt_surname or flt_surname = '')
    and (p.patronymic = flt_patronymic or flt_patronymic = '')
    and (p.gender = flt_gender or flt_gender = '')
    and (p.age = flt_age or flt_age = 0)
    and (p.nationality = flt_nationality or flt_nationality = '')
  order by id
  limit per_page+1 offset (page_id-1)*per_page
  );
END;
$$ LANGUAGE plpgsql;
