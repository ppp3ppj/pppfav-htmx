BEGIN;

-- Check and drop foreign key constraints if they exist
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'member_person_id_fkey'
        AND table_name = 'Member'
    ) THEN
        ALTER TABLE "Member" DROP CONSTRAINT "member_person_id_fkey";
    END IF;

    IF EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'member_girl_group_id_fkey'
        AND table_name = 'Member'
    ) THEN
        ALTER TABLE "Member" DROP CONSTRAINT "member_girl_group_id_fkey";
    END IF;

    IF EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'song_girl_group_id_fkey'
        AND table_name = 'Song'
    ) THEN
        ALTER TABLE "Song" DROP CONSTRAINT "song_girl_group_id_fkey";
    END IF;
END $$;

-- Drop tables if they exist
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'Song') THEN
        DROP TABLE "Song";
    END IF;

    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'Member') THEN
        DROP TABLE "Member";
    END IF;

    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'GirlGroup') THEN
        DROP TABLE "GirlGroup";
    END IF;

    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'Person') THEN
        DROP TABLE "Person";
    END IF;
END $$;

-- Drop sequences if they exist
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_class WHERE relname = 'person_id_seq') THEN
        DROP SEQUENCE person_id_seq;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_class WHERE relname = 'girlgroup_id_seq') THEN
        DROP SEQUENCE girlgroup_id_seq;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_class WHERE relname = 'member_id_seq') THEN
        DROP SEQUENCE member_id_seq;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_class WHERE relname = 'song_id_seq') THEN
        DROP SEQUENCE song_id_seq;
    END IF;
END $$;

-- Drop the extension safely
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'uuid-ossp') THEN
        DROP EXTENSION "uuid-ossp";
    END IF;
END $$;

COMMIT;
