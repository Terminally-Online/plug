-- AlterTable
CREATE SEQUENCE socketidentity_onboardingcount_seq;
ALTER TABLE "SocketIdentity" ALTER COLUMN "onboardingCount" SET DEFAULT nextval('socketidentity_onboardingcount_seq');
ALTER SEQUENCE socketidentity_onboardingcount_seq OWNED BY "SocketIdentity"."onboardingCount";
