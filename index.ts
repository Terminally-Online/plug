import sharp from "sharp";
import path from "path";

import { default as fse } from "fs-extra";
import { createHash } from "crypto";
import { promises as fs } from "fs";
import { PDFDocument, PDFString, PDFName } from "pdf-lib";
import { exec } from "child_process";

const rename = async (oldPath: string, newPath: string) => {
    await fse.rename(oldPath, newPath);
};

const hashFile = async (filePath: string) => {
    const fileBuffer = await fs.readFile(filePath);
    const hashSum = createHash("sha256");
    hashSum.update(fileBuffer);
    return hashSum.digest("hex");
};

const addWatermark = async (
    imagePath: string,
    email: string,
    tag: string,
    finalDeckHash: string
) => {
    const image = sharp(imagePath);
    const metadata = await image.metadata();

    // @ts-ignore
    const fontSize = Math.floor(metadata?.width / 20);
    const svgImage = Buffer.from(`
        <svg width="${metadata.width}" height="${metadata.height}">
            <text x="50%" y="50%" alignment-baseline="middle" text-anchor="middle" font-family="Arial" font-size="${fontSize}px" fill="white" opacity="0.3" transform="rotate(-45 ${
                // @ts-ignore
                metadata?.width / 2
                // @ts-ignore
            }, ${metadata.height / 2})">
                ${email}
            </text>
        </svg>
    `);

    const outputPath = `./outputs/${
        email !== "" ? email : "default"
    }/${tag}/${finalDeckHash}/${path.basename(imagePath)}`;
    await fse.ensureDir(path.dirname(outputPath));
    await image
        .composite([{ input: svgImage, blend: "over" }])
        .toFile(outputPath);
};

const createPDF = async (imagePaths: string[], outputPath: string) => {
    const pdfDoc = await PDFDocument.create();
    for (const imagePath of imagePaths) {
        const imageBytes = await fs.readFile(imagePath);
        const image = await pdfDoc.embedPng(imageBytes);
        const page = pdfDoc.addPage([image.width, image.height]);
        page.drawImage(image, {
            x: 0,
            y: 0,
            width: image.width,
            height: image.height,
        });
    }
    const pdfBytes = await pdfDoc.save();
    await fs.writeFile(outputPath, pdfBytes);
};

const vision = async (email: string) => {
    const files = (await fse.readdir("./base/vision"))
        .filter((file) => file.endsWith(".png"))
        .sort((a, b) => {
            const matchA = a.match(/\d+/);
            const matchB = b.match(/\d+/);

            const numA = matchA ? parseInt(matchA[0], 10) : 0;
            const numB = matchB ? parseInt(matchB[0], 10) : 0;

            return numA - numB;
        });

    // Truncate names to numbers and hash the image contents
    let finalHashData = "";
    for (const file of files) {
        const newName = file.replace("Seed - ", "");
        const oldPath = `./base/vision/${file}`;
        const newPath = `./base/vision/${newName}`;
        await rename(oldPath, newPath);
        const hash = await hashFile(newPath);
        finalHashData += hash;
    }

    // Calculate the final hash of authenticity
    const finalHashSum = createHash("sha256");
    finalHashSum.update(finalHashData);
    const finalDeckHash = "0x" + finalHashSum.digest("hex").slice(0, 16);

    // Watermark it when it is going to a specific person
    for (const file of files) {
        const newPath = `./base/vision/${file.replace("Seed - ", "")}`;
        await addWatermark(
            newPath,
            email !== "" ? email : "",
            "vision-images",
            finalDeckHash
        );
    }

    // Create the pdf of the watermarked deck images
    await createPDF(
        files.map(
            (file) =>
                `./outputs/${
                    email !== "" ? email : "default"
                }/vision-images/${finalDeckHash}/${file.replace("Seed - ", "")}`
        ),
        `./outputs/${email !== "" ? email : "default"}/vision.pdf`
    );

    const absolutePath = path.resolve(
        `./outputs/${email !== "" ? email : "default"}/vision.pdf`
    );

    // use ffmpeg to create a gif from the images
    exec(
        `ffmpeg -y -framerate 1 -start_number 1 -i './outputs/${
            email !== "" ? email : "default"
        }/vision-images/${finalDeckHash}/%d.png' -vf "palettegen=stats_mode=full" './outputs/${
            email !== "" ? email : "default"
        }/vision-images/${finalDeckHash}/palette.png' && ffmpeg -y -framerate 1 -start_number 1 -i './outputs/${
            email !== "" ? email : "default"
        }/vision-images/${finalDeckHash}/%d.png' -i './outputs/${
            email !== "" ? email : "default"
        }/vision-images/${finalDeckHash}/palette.png' -filter_complex "[0:v]fps=1,setpts=6.0*PTS[v];[v][1:v]paletteuse=dither=bayer:bayer_scale=5" './outputs/${
            email !== "" ? email : "default"
        }/vision.gif'`,
        (error, stdout, stderr) => {
            if (error) {
                console.error(`exec error: ${error}`);
                return;
            }
            console.log(`stdout: ${stdout}`);
            console.error(`stderr: ${stderr}`);
        }
    );

    console.log(`Vision deck generated:
    for: ${email}
    path (cmd + click to open): ${absolutePath} 
    hash: ${finalDeckHash}`);
};

const boomer = async (email: string) => {
    const files = (await fse.readdir("./base/boomer"))
        .filter((file) => file.endsWith(".png"))
        .sort((a, b) => {
            const matchA = a.match(/\d+/);
            const matchB = b.match(/\d+/);

            const numA = matchA ? parseInt(matchA[0], 10) : 0;
            const numB = matchB ? parseInt(matchB[0], 10) : 0;

            return numA - numB;
        });

    // Truncate names to numbers and hash the image contents
    let finalHashData = "";
    for (const file of files) {
        const newName = file.replace("Seed - Boomer - ", "");
        const oldPath = `./base/boomer/${file}`;
        const newPath = `./base/boomer/${newName}`;
        await rename(oldPath, newPath);
        const hash = await hashFile(newPath);
        finalHashData += hash;
    }

    // Calculate the final hash of authenticity
    const finalHashSum = createHash("sha256");
    finalHashSum.update(finalHashData);
    const finalDeckHash = "0x" + finalHashSum.digest("hex").slice(0, 16);

    // Watermark it when it is going to a specific person
    for (const file of files) {
        const newPath = `./base/boomer/${file.replace("Seed - Boomer - ", "")}`;
        await addWatermark(
            newPath,
            email !== "" ? email : "",
            "boomer-images",
            finalDeckHash
        );
    }

    // Create the pdf of the watermarked deck images
    await createPDF(
        files.map(
            (file) =>
                `./outputs/${
                    email !== "" ? email : "default"
                }/boomer-images/${finalDeckHash}/${file.replace(
                    "Seed - Boomer - ",
                    ""
                )}`
        ),
        `./outputs/${email !== "" ? email : "default"}/deck.pdf`
    );

    const absolutePath = path.resolve(
        `./outputs/${email !== "" ? email : "default"}/vision.pdf`
    );

    // use ffmpeg to create a gif from the images
    exec(
        `ffmpeg -y -framerate 1 -start_number 2 -i './outputs/${
            email !== "" ? email : "default"
        }/boomer-images/${finalDeckHash}/%d.png' -vf "palettegen=stats_mode=full" './outputs/${
            email !== "" ? email : "default"
        }/boomer-images/${finalDeckHash}/palette.png' && ffmpeg -y -framerate 1 -start_number 2 -i './outputs/${
            email !== "" ? email : "default"
        }/boomer-images/${finalDeckHash}/%d.png' -i './outputs/${
            email !== "" ? email : "default"
        }/boomer-images/${finalDeckHash}/palette.png' -filter_complex "[0:v]fps=1,setpts=16.0*PTS[v];[v][1:v]paletteuse=dither=bayer:bayer_scale=5" './outputs/${
            email !== "" ? email : "default"
        }/deck.gif'`,
        (error, stdout, stderr) => {
            if (error) {
                console.error(`exec error: ${error}`);
                return;
            }
            console.log(`stdout: ${stdout}`);
            console.error(`stderr: ${stderr}`);
        }
    );

    console.log(`Deck generated:
    for: ${email}
    path (cmd + click to open): ${absolutePath}
    hash: ${finalDeckHash}`);
};

const onePager = async (email: string) => {
    const links = {
        chance: "https://calendly.com/chance-zyez/chance",
        reka: "https://calendly.com/d/cp33-4v2-qn9/reka-chance",
        drake: "https://calendly.com/d/ck7n-gng-zwv/chance-drake",
    };

    for (const [name, link] of Object.entries(links)) {
        const imagePath = "./base/one-pager.png";
        const pdfDoc = await PDFDocument.create();
        const imageBytes = await fs.readFile(imagePath);
        const image = await pdfDoc.embedPng(imageBytes);
        const page = pdfDoc.addPage([image.width, image.height]);
        page.drawImage(image, {
            x: 0,
            y: 0,
            width: image.width,
            height: image.height,
        });

        const annotation = page.doc.context.register(
            page.doc.context.obj({
                Type: "Annot",
                Subtype: "Link",
                Rect: [0, 0, image.width, 41 * 4],
                C: [0, 0, 1],
                A: {
                    Type: "Action",
                    S: "URI",
                    URI: PDFString.of(link),
                },
            })
        );
        page.node.set(PDFName.of("Annots"), pdfDoc.context.obj([annotation]));

        const pdfBytes = await pdfDoc.save();
        await fse.ensureDir(
            path.dirname(`./outputs/${email !== "" ? email : "default"}/`)
        );
        await fs.writeFile(
            `./outputs/${
                email !== "" ? email : "default"
            }/plug-with-${name}-one-pager.pdf`,
            pdfBytes
        );
    }

    console.log(`One pager generated:
    for: ${email}
    path (cmd + click to open): ${path.resolve(
        `./outputs/${email !== "" ? email : "default"}/one-pager.pdf`
    )}`);
};

const memo = async (email: string) => {
    const files = (await fse.readdir("./base/memo"))
        .filter((file) => file.endsWith(".png"))
        .sort((a, b) => {
            const matchA = a.match(/\d+/);
            const matchB = b.match(/\d+/);

            const numA = matchA ? parseInt(matchA[0], 10) : 0;
            const numB = matchB ? parseInt(matchB[0], 10) : 0;

            return numA - numB;
        });

    // Truncate names to numbers and hash the image contents
    let finalHashData = "";
    for (const file of files) {
        const newName = file.replace("Investor Memo - ", "");
        const oldPath = `./base/memo/${file}`;
        const newPath = `./base/memo/${newName}`;
        await rename(oldPath, newPath);
        const hash = await hashFile(newPath);
        finalHashData += hash;
    }

    // Calculate the final hash of authenticity
    const finalHashSum = createHash("sha256");
    finalHashSum.update(finalHashData);
    const finalDeckHash = "0x" + finalHashSum.digest("hex").slice(0, 16);

    // Watermark it when it is going to a specific person
    for (const file of files) {
        const newPath = `./base/memo/${file.replace("Investor Memo - ", "")}`;
        await addWatermark(
            newPath,
            email !== "" ? email : "",
            "memo-images",
            finalDeckHash
        );
    }

    // Create the pdf of the watermarked deck images
    await createPDF(
        files.map(
            (file) =>
                `./outputs/${
                    email !== "" ? email : "default"
                }/memo-images/${finalDeckHash}/${file}`
        ),
        `./outputs/${email !== "" ? email : "default"}/memo.pdf`
    );

    const absolutePath = path.resolve(
        `./outputs/${email !== "" ? email : "default"}/memo.pdf`
    );

    console.log(`Memo generated:
    for: ${email}
    path (cmd + click to open): ${absolutePath}
    hash: ${finalDeckHash}`);
};

const main = async (email: string) => {
    await vision(email);
    await boomer(email);
    await onePager(email);
    await memo(email);
};

const email = process.argv[2] ?? "";

main(email).then(() => console.log("✔︎ Generated raise components"));
