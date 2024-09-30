import { join } from "path";
import {
    readFileSync,
    readdirSync,
    readJSONSync,
    writeFileSync,
} from "fs-extra";

const main = async () => {
    const templatesDir = join(__dirname, "..", "templates");
    const templates = readdirSync(templatesDir);
    const templateNames = templates.map((template) => {
        const split = template.split(".");
        split.pop();
        return split.join(".");
    });

    const packagesDir = join(__dirname, "..", "..", "..", "packages");
    const packages = readdirSync(packagesDir);
    const packageNames = packages.filter(
        (packageName) => !templateNames.includes(packageName)
    );

    packageNames.forEach((packageName) => {
        if (packageName.includes(".")) return;

        const packageDir = join(packagesDir, packageName);
        const packageJson = join(packageDir, "package.json");
        const hasPackageJson = readdirSync(packageDir).includes("package.json");

        if (!hasPackageJson) return;

        let { license } = readJSONSync(packageJson);
        const lowerLicense = license ? license.toLowerCase() : "mit";
        const { name, version, ...rest } = readJSONSync(packageJson);

        if (templateNames.includes(lowerLicense)) {
            writeFileSync(
                packageJson,
                JSON.stringify(
                    {
                        name,
                        version,
                        license,
                        ...rest,
                    },
                    null,
                    4
                )
            );

            writeFileSync(
                join(packageDir, "LICENSE"),
                readFileSync(join(templatesDir, `${lowerLicense}.txt`))
                    .toString()
                    .replace(
                        /<PACKAGE_NAME>/g,
                        `@terminally-online/plug-${packageName}`
                    )
            );

            console.log(
                `✔︎ Applied ${lowerLicense} to /packages/${packageName}/`
            );
        } else {
            console.error(
                `⨯ Unknown license: ${lowerLicense}. Available options include:\n   - ${templateNames.join(
                    "\n   - "
                )}`
            );
        }
    });
};

main();
