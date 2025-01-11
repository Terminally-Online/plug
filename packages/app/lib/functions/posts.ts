
        export const PAGE_SIZE = 20;

        export interface Post {
            filename: string;
            slug: string;
            title: string;
            description: string;
            image: string;
            content: string;
            attributes: {
                created: string;
            } & Partial<{
                updated: string;
                tags: string[];
                related: string[];
                inbound: string[];
                author: string;
                // display settings
                imagePosition: 'top' | 'bottom';
                imagePadded: 'true' | 'false';
                className: string;
                variant: string;
                unlisted: 'true' | 'false';
                sidebar: 'show' | 'hide' | undefined;
            }>
        }

        export type Posts = Record<string, Post>;

        export const faviconUrls = {"docs.onplug.io":"data:image/vnd.microsoft.icon;base64,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SP8Z503/GedN/xfmU/8X5lP/F+ZT/xXmWf8V5ln/E+Vf/xPlX/8T5V//EORl/xDkZf8Q5GX/DuRr/w7ka/8L43H/C+Nx/wvjcf8K4nj/CuJ4/zDuDv8w7g7/MO4O/y3tFv8t7Rb/Le0W/yrsHf8q7B3/Kuwd/ybrI/8m6yP/Jusj/yXrKv8l6yr/Jesq/yPqMv8j6jL/3fzk///////5//r//////2XweP8b6Tn/HuhA/x7oQP8b6Ej/G+hI/xnnTf8Z503/GedN/xfmU/8X5lP/F+ZT/xXmWf8V5ln/E+Vf/xPlX/8T5V//EORl/xDkZf8O5Gv/DuRr/w7ka/8L43H/C+Nx/wrieP8K4nj/CuJ4/zDuDv8w7g7/MO4O/y3tFv8t7Rb/Le0W/yrsHf8q7B3/Kuwd/ybrI/8m6yP/Jusj/yXrKv8l6yr/Jesq/yPqMv8j6jL/3fzk///////5//r//////2XweP8b6Tn/HuhA/xvoSP8b6Ej/G+hI/xnnTf8Z503/F+ZT/xfmU/8X5lP/FeZZ/xXmWf8T5V//E+Vf/xPlX/8Q5GX/EORl/xDkZf8O5Gv/DuRr/wvjcf8L43H/C+Nx/wrieP8K4nj/CuJ4/zDuDv8w7g7/MO4O/y3tFv8t7Rb/Le0W/yrsHf8q7B3/Kuwd/ybrI/8m6yP/Jesq/yXrKv8l6yr/Jesq/yPqMv8j6jL/3fzk///////5//r//////2XweP8b6Tn/G+hI/xvoSP8b6Ej/G+hI/xnnTf8X5lP/F+ZT/xfmU/8X5lP/FeZZ/xXmWf8T5V//E+Vf/xPlX/8Q5GX/EORl/xDkZf8O5Gv/DuRr/wvjcf8L43H/C+Nx/wrieP8K4nj/CuJ4/zDuDv8w7g7/Le0W/y3tFv8t7Rb/Kuwd/yrsHf8q7B3/Jusj/ybrI/8m6yP/Jesq/yXrKv8l6yr/I+oy/yPqMv8j6jL/3fzk///////5//r//////2XweP8b6Tn/G+hI/xvoSP8b6Ej/G+hI/xnnTf8F5E7/BeRO/xfmU/8V5ln/FeZZ/xXmWf8T5V//E+Vf/xPlX/8Q5GX/EORl/w7ka/8O5Gv/DuRr/wvjcf8L43H/CuJ4/wrieP8K4nj/B+J//zDuDv8w7g7/Le0W/y3tFv8t7Rb/Kuwd/yrsHf8q7B3/Jusj/ybrI/8m6yP/Jesq/yXrKv8l6yr/I+oy/yPqMv8j6jL/3fzk///////5//r//////2XweP8b6Tn/vfjO/9385P/d/OT/3fzk/9385P+9+M7/ffGk/xPlX/8F5E7/FeZZ/xXmWf8T5V//E+Vf/xDkZf8Q5GX/EORl/w7ka/8O5Gv/C+Nx/wvjcf8L43H/CuJ4/wrieP8K4nj/B+J//zDuDv8w7g7/Le0W/y3tFv8t7Rb/Kuwd/yrsHf8q7B3/Jusj/ybrI/8l6yr/Jesq/yXrKv8l6yr/Jesq/yPqMv8j6jL/3fzk///////5//r//////2rwgf8b6Tn/nPWu///////5//r///////////////////////n/+v998aT/BeRO/xPlX/8T5V//E+Vf/xDkZf8Q5GX/EORl/w7ka/8O5Gv/C+Nx/wvjcf8L43H/CuJ4/wrieP8K4nj/B+J//zDuDv8t7Rb/Le0W/y3tFv8q7B3/Kuwd/yrsHf8m6yP/Jusj/ybrI/8m6yP/Jesq/zbsSP827Ej/NuxI/zbsSP8j6jL/3fzk///////5//r//////2rwgf8b6Tn/O+pw///////////////////////5//r/+f/6////////////ffGk/wXkTv8T5V//E+Vf/xDkZf8Q5GX/DuRr/w7ka/8O5Gv/C+Nx/wvjcf8L43H/CuJ4/wrieP8K4nj/B+J//zDuDv8t7Rb/Le0W/y3tFv8q7B3/Kuwd/yrsHf8m6yP/Jusj/ybrI/8j6jL/vfjO//n/+v/5//r/+f/6/9385P827Ej/3fzk///////5//r//////2rwgf8b6Tn/G+hI/734zv/d/OT/3fzk//n/+v/////////////////5//r//////0zrkv8F5E7/EORl/xDkZf8Q5GX/DuRr/w7ka/8L43H/C+Nx/wvjcf8K4nj/CuJ4/wrieP8H4n//B+J//zDuDv8t7Rb/Le0W/y3tFv8q7B3/Kuwd/yrsHf8m6yP/Jesq/ybrI/9q8IH///////n/+v////////////n/+v827Ej/3fzk///////5//r//////2rwgf8b6Tn/G+hI/xvoSP8Z503/F+ZT/zvqcP+c9a7//////////////////////9385P8T5V//EORl/xDkZf8O5Gv/DuRr/w7ka/8L43H/C+Nx/wvjcf8K4nj/CuJ4/wrieP8H4n//B+J//y3tFv8t7Rb/Le0W/yrsHf8q7B3/Kuwd/yrsHf8m6yP/Jesq/ybrI/9l8Hj///////////////////////n/+v827Ej/3fzk///////5//r//////2rwgf8b6Tn/G+hI/xvoSP8Z503/GedN/xnnTf8F5E7/nPWu//////////////////////876nD/BeRO/xDkZf8Q5GX/EORl/xDkZf8L43H/C+Nx/wvjcf8K4nj/CuJ4/wrieP8H4n//B+J//y3tFv8t7Rb/Le0W/yrsHf8q7B3/Kuwd/ybrI/8m6yP/Jusj/yXrKv8m6yP/ZfB4/5z1rv+c9a7/nPWu/5z1rv827Ej/3fzk///////5//r//////2rwgf8b6Tn/G+hI/xnnTf8Z503/GedN/xfmU/8X5lP/O+pw//n/+v/////////////////d/OT/vfjO/734zv+9+M7/vfjO/734zv9M65L/C+Nx/wrieP8K4nj/CuJ4/wfif/8H4n//B+J//y3tFv8t7Rb/Le0W/yrsHf8q7B3/Kuwd/ybrI/8m6yP/Jesq/yXrKv8l6yr/Jusj/ybrI/8m6yP/Jusj/ybrI/8b6Tn/3fzk///////5//r//////2rwgf8b6Tn/GedN/xnnTf8Z503/F+ZT/xfmU/8X5lP/FeZZ/9385P////////////////////////////////////////////////9M65L/C+Nx/wrieP8K4nj/CuJ4/wfif/8H4n//B+J//y3tFv8t7Rb/Le0W/yrsHf8q7B3/Kuwd/ybrI/8m6yP/Jesq/yXrKv8l6yr/ZfB4/5z1rv+c9a7/nPWu/5z1rv827Ej/3fzk///////5//r//////2rwgf8b6Tn/GedN/xnnTf8Z503/F+ZT/xfmU/8X5lP/FeZZ/9385P////////////////////////////////////////////////9M65L/C+Nx/wrieP8K4nj/CuJ4/wfif/8H4n//B+J//y3tFv8t7Rb/Kuwd/yrsHf8q7B3/Jusj/ybrI/8m6yP/Jesq/ybrI/9l8Hj///////////////////////n/+v827Ej/3fzk///////5//r//////2rwgf8b6Tn/GedN/xnnTf8Z503/F+ZT/xXmWf8X5lP/E+Vf//n/+v/////////////////d/OT/3fzk/9385P/d/OT/3fzk/9385P9M65L/C+Nx/wrieP8K4nj/B+J//wfif/8H4n//BOGG/i3tFv8t7Rb/Kuwd/yrsHf8q7B3/Jusj/ybrI/8l6yr/Jesq/ybrI/9q8IH///////n/+v////////////n/+v827Ej/3fzk///////5//r//////2rwgf8b6Tn/GedN/xnnTf8X5lP/F+ZT/xfmU/8F5E7/nPWu//////////////////////9M65L/EORl/w7ka/8O5Gv/C+Nx/wvjcf8L43H/CuJ4/wrieP8K4nj/B+J//wfif/8H4n//BOGG/i3tFv8t7Rb/Kuwd/yrsHf8q7B3/Jusj/ybrI/8l6yr/Jesq/yXrKv827Ej/vfjO//n/+v/5//r/+f/6/9385P827Ej/3fzk///////5//r//////2rwgf8b6Tn/GedN/xnnTf8X5lP/F+ZT/zvqcP+c9a7//////////////////////9385P8Q5GX/EORl/w7ka/8O5Gv/C+Nx/wvjcf8L43H/CuJ4/wrieP8K4nj/B+J//wfif/8H4n//BOGG/i3tFv8q7B3/Kuwd/yrsHf8m6yP/Jusj/ybrI/8l6yr/Jesq/yXrKv8l6yr/I+oy/zbsSP827Ej/NuxI/zbsSP8e6ED/3fzk///////5//r//////2XweP+c9a7/3fzk/9385P/d/OT/3fzk//n/+v/////////////////5//r//////0zrkv8T5V//DuRr/w7ka/8O5Gv/C+Nx/wvjcf8K4nj/CuJ4/wrieP8H4n//B+J//wfif/8E4Yb+BOGG/i3tFv8q7B3/Kuwd/yrsHf8m6yP/Jusj/ybrI/8l6yr/Jesq/yXrKv8j6jL/I+oy/yPqMv8j6jL/G+k5/xvpOf8e6ED/3fzk///////5//r//////2XweP+9+M7////////////////////////////5//r/+f/6////////////nPWu/xPlX/8O5Gv/DuRr/w7ka/8L43H/C+Nx/wvjcf8K4nj/CuJ4/wrieP8H4n//B+J//wfif/8E4Yb+BOGG/i3tFv8q7B3/Kuwd/yrsHf8m6yP/Jusj/yXrKv8l6yr/Jesq/yXrKv8j6jL/I+oy/yPqMv8j6jL/G+k5/x7o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I/+e9qv///////////9P7mv/Eec7/xzoRv8c6Eb/GedO/xnnTv8Z507/F+dV/xPlWv8T5Vr/E+Vf/xDkY/8P5Gj/D+Ro/w3jbf8M43L/CuJ2/wridv8v7g//L+4P/y/uD/8s7Rn/LO0Z/ynsIf8p7CH/Juso/ybrKP8j6jH/FOkj/572q////////////0/ua/8R5zv/HOhG/xzoRv8Z507/GedO/xfnVf8X51X/E+Va/xPlX/8T5V//EORj/w/kaP8N423/DeNt/wzjcv8K4nb/COJ7/y/uD/8v7g//L+4P/yztGf8s7Rn/Kewh/ynsIf8m6yj/Juso/yPqMf8U6SP/nvar////////////T+5r/xHnO/8c6Eb/HOhG/xnnTv8Z507/F+dV/xfnVf8T5Vr/E+Vf/xPlX/8Q5GP/D+Ro/w3jbf8N423/DONy/wridv8I4nv/L+4P/y/uD/8s7Rn/LO0Z/yztGf8p7CH/Kewh/ybrKP8m6yj/I+ox/xTpI/+e9qv///////////9P7mv/Eec7/xzoRv8Z507/GedO/wfkS/8H5Ev/E+Va/xPlWv8T5V//EORj/w/kaP8P5Gj/DeNt/wzjcv8K4nb/CuJ2/wjie/8v7g//L+4P/yztGf8s7Rn/Kewh/ynsIf8p7CH/Juso/ybrKP8j6jH/FOkj/572q////////////0/ua/9v8YT/9P73/+D85v/O+t3/pfa+/z7qev8H5Ev/E+Vf/xPlX/8Q5GP/D+Ro/w/kaP8N423/DONy/wridv8K4nb/COJ7/y/uD/8v7g//LO0Z/yztGf8p7CH/Kewh/ynsIf8m6yj/OuxJ/zrsSf8j6jH/nvar////////////T+5r/0/ua////////////////////////////572q/8T5Vr/EORj/xDkY/8P5Gj/DeNt/w3jbf8M43L/CuJ2/wjie/8I4nv/L+4P/yztGf8s7Rn/LO0Z/ynsIf8p7CH/Juso/6X2vv/0/vf//////572q/+e9qv///////////9P7mv/Eec7/572q/+l9r7/zvrd/////////////////2/xhP8T5Vr/D+Ro/w/kaP8N423/DONy/wzjcv8K4nb/COJ7/wfhgP8v7g//LO0Z/yztGf8s7Rn/Kewh/ynsIf8m6yj/4Pzm////////////nvar/572q////////////0/ua/8R5zv/Eec7/wfkS/8H5Ev/nvar////////////zvrd/xPlWv8Q5GP/D+Ro/w3jbf8M43L/CuJ2/wridv8I4nv/B+GA/y/uD/8s7Rn/LO0Z/ynsIf8p7CH/Juso/ybrKP867En/b/GE/2/xhP867En/nvar////////////T+5r/xHnO/8Z507/F+dV/xnnTv8T5V//9P73///////0/vf/zvrd/8763f/O+t3/pfa+/wzjcv8K4nb/COJ7/wjie/8H4YD/LO0Z/yztGf8s7Rn/Kewh/ynsIf8m6yj/Juso/zrsSf9P7mv/b/GE/zrsSf+e9qv///////////9P7mv/Eec7/xnnTv8Z507/F+dV/xPlWv/g/Ob////////////////////////////O+t3/DONy/wridv8I4nv/B+GA/wfhgP8s7Rn/LO0Z/yztGf8p7CH/Kewh/ybrKP8j6jH/4Pzm////////////nvar/572q////////////0/ua/8R5zv/GedO/xfnVf8H5Ev/Pup6////////////9P73/6X2vv+l9r7/pfa+/3Pwn/8M43L/CuJ2/wjie/8H4YD/B+GA/yztGf8s7Rn/Kewh/ynsIf8m6yj/Juso/ybrKP+l9r7/9P73//////+e9qv/nvar////////////T+5r/xzoRv8X51X/F+dV/z7qev/O+t3///////////+l9r7/B+RL/xDkY/8Q5GP/DeNt/wridv8I4nv/COJ7/wfhgP8E4Yf+LO0Z/yztGf8p7CH/Kewh/ybrKP8m6yj/Juso/yPqMf867En/OuxJ/x/pO/+e9qv///////T+9/9v8YT/zvrd/+D85v/g/Ob/9P73////////////9P73/z7qev8P5Gj/DeNt/wzjcv8K4nb/CuJ2/wjie/8H4YD/B+GA/wThh/4s7Rn/LO0Z/ynsIf8p7CH/Juso/ybrKP8j6jH/I+ox/yPqMf8f6Tv/FOkj/572q///////9P73/3Pwn//0/vf//////////////////////+D85v8+6nr/EORj/w3jbf8N423/DONy/wridv8K4nb/COJ7/wfhgP8H4YD/BOGH/iztGf8p7CH/Kewh/ybrKP8m6yj/Juso/yPqMf8j6jH/H+k7/x/pO/8R5zv/b/GE/6X2vv+e9qv/T+5r/572q/+l9r7/pfa+/572q/9z8J//EORj/xPlWv8P5Gj/DeNt/w3jbf8M43L/CuJ2/wjie/8I4nv/B+GA/wfhgP8E4Yf+LO0Z/ynsIf8p7CH/Juso/ybrKP8j6jH/I+ox/yPqMf8f6Tv/H+k7/x/pO/8R5zv/Eec7/xHnO/8Z507/B+RL/wfkS/8H5Ev/B+RL/wfkS/8Q5GP/D+Ro/w/kaP8N423/DONy/wzjcv8K4nb/COJ7/wfhgP8H4YD/BOGH/gThh/4p7CH/Kewh/ynsIf8m6yj/Juso/yPqMf8j6jH/I+ox/x/pO/8f6Tv/H+k7/xzoRv8c6Eb/GedO/xnnTv8X51X/F+dV/xPlWv8T5V//E+Vf/xDkY/8P5Gj/D+Ro/w3jbf8M43L/CuJ2/wridv8I4nv/B+GA/wfhgP8E4Yf+BOGH/i3tFfUp7CH/Juso/ybrKP8m6yj/I+ox/yPqMf8j6jH/H+k7/x/pO/8c6Eb/HOhG/xzoRv8Z507/GedO/xfnVf8X51X/E+Va/xPlX/8T5V//EORj/w/kaP8N423/DeNt/wzjcv8K4nb/COJ7/wjie/8H4YD/B+GA/wThh/4E4Yf+K+wd4ynsIf8m6yj/Juso/yPqMf8j6jH/I+ox/x/pO/8f6Tv/H+k7/xzoRv8c6Eb/HOhG/xnnTv8Z507/F+dV/xfnVf8T5Vr/E+Vf/xPlX/8Q5GP/D+Ro/w3jbf8N423/DONy/wridv8I4nv/COJ7/wfhgP8E4Yf+BOGH/grid+Mr7Ru3Kewh/ybrKP8m6yj/I+ox/yPqMf8j6jH/H+k7/x/pO/8f6Tv/HOhG/xzoRv8Z507/GedO/xnnTv8X51X/E+Va/xPlWv8T5V//EORj/w/kaP8P5Gj/DeNt/wzjcv8K4nb/CuJ2/wjie/8H4YD/B+GA/wThh/4E4Yf+COJ9qy7tElMm6yj/Juso/ybrKP8j6jH/I+ox/yPqMf8f6Tv/H+k7/xzoRv8c6Eb/HOhG/xnnTv8Z507/F+dV/xfnVf8T5Vr/E+Vf/xPlX/8Q5GP/D+Ro/w3jbf8N423/DONy/wridv8I4nv/COJ7/wfhgP8H4YD/BOGH/gThh/4J4npSHOhOASztGqIm6yj/I+ox/yPqMf8j6jH/H+k7/x/pO/8f6Tv/HOhG/xzoRv8c6Eb/GedO/xnnTv8X51X/F+dV/xPlWv8T5V//E+Vf/xDkY/8P5Gj/DeNt/w3jbf8M43L/CuJ2/wjie/8I4nv/B+GA/wThh/4E4Yf+COJ9qxzoTgEc6E4BHOhOASztGqIj6jH/I+ox/yPqMf8f6Tv/H+k7/x/pO/8c6Eb/HOhG/xzoRv8Z507/GedO/xfnVf8T5Vr/E+Va/xPlX/8Q5GP/D+Ro/w/kaP8N423/DONy/wridv8K4nb/COJ7/wfhgP8H4YD/BOGH/gjifasc6E4BHOhOARzoTgEc6E4BHOhOASTqL1Er7Ru3K+wd4x/pO/8f6Tv/HOhG/xzoRv8c6Eb/GedO/xnnTv8X51X/F+dV/xPlWv8T5Vr/E+Vf/xDkY/8P5Gj/D+Ro/w3jbf8M43L/CuJ2/wridv8I4nv/CuJ34wjifasJ4npSHOhOARzoTgEc6E4B8AAAD8AAAAOAAAABgAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAGAAAABwAAAA/AAAA8oAAAAEAAAACAAAAABACAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAKehyASztGoIu7RXqLO0Z/ynsIf8n6yn/I+oy/yPqMv8h6Tv/HOhI/xzoSP8Y50//E+ZW/w/kY/4O5GqCKehyASztGoIv7hD/LO0Z/yztGf8p7CH/J+sp/yPqMv8h6Tv/Gug//xzoSP8Y50//E+ZW/xTmXP8P5GP+DuRq/w7kaoIu7RXqL+4Q/yztGf8p7CH/J+sp/yfrKf8Z6h3/Iek7/xroP/8c6Ej/GOdP/xPmVv8U5lz/D+Rj/g7kav8M43HnL+4Q/y/uEP8s7Rn/Kewh/xnqHf9T71r/v/jN/znsUP8a6D//HOhI/xjnT/8T5lb/FOZc/w/kY/4O5Gr/CuJ2/y/uEP8v7hD/LO0Z/ynsIf8Z6h3/afF3//////857FD/Gug//xjnT/8T5lb/FOZc/w/kY/4O5Gr/C+Nx/wridv8v7hD/LO0Z/ynsIf8n6yn/Geod/2nxd///////Puti/xzoSP8Y50//CeVE/xTmXP8P5GP+DuRq/wvjcf8K4nb/L+4Q/yztGf8p7CH/J+sp/yPqMv9p8Xf//////2nxd//0/vf/9P73/3/yk/8T5lb/D+Rj/g7kav8K4nb/COJ8/y/uEP8s7Rn/Geod/2nxd//0/vf/qPa2//////8+62L/Tux7/6j2tv//////Vu2P/w/kY/4L43H/CuJ2/wjifP8s7Rn/Kewh/ynsIf857FD/f/KT/3/yk///////Puti/wnlRP8Y50//9P73//T+9//0/vf/aO6c/wvjcf8G4YT/LO0Z/ynsIf8Z6h3/afF3//T+9/+o9rb/9P73/z7rYv8Y50//Vu2P//////+U87v/Vu2P/y7niP8K4nb/BuGE/yztGf8p7CH/J+sp/yPqMv8h6Tv/afF3//T+9/+/+M3///////////+/+M3/D+Rj/g/kY/4L43H/COJ8/wbhhP8p7CH/Kewh/yfrKf8j6jL/I+oy/znsUP9p8Xf/Tux7/2junP9O7Hv/D+Rj/g7kav8L43H/CuJ2/wjifP8G4YT/Kewh/yfrKf8n6yn/I+oy/yHpO/8a6D//CeVE/wnlRP8J5UT/E+ZW/w/kY/4O5Gr/C+Nx/wjifP8G4YT/BuGE/y7tFeon6yn/I+oy/yPqMv8h6Tv/HOhI/xjnT/8Y50//FOZc/w/kY/4O5Gr/C+Nx/wridv8I4nz/BuGE/wXhhews7RqCJ+sp/yPqMv8h6Tv/Gug//xzoSP8Y50//E+ZW/xTmXP8P5GP+DuRq/wvjcf8K4nb/COJ8/wbhhP8E4ImCKehyASztGoIj6jPwIek7/xroP/8c6Ej/GOdP/xPmVv8U5lz/D+Rj/g7kav8L43H/COJ8/wXhhewE4ImCKehyAYABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIABAAA="} as const;

        export const posts: Posts = {helloworld:{filename:"say-hello",slug:"hello-world",title:"Hello World.",description:"In a market that never sleeps, Plug is your heartbeat that keeps your onchain presence constant....",image:"/cdn/papers/hello-world.png",content:"\n\nTLDR: After a year of building in stealth, Plug is ready for you to try. We're making onchain automation accessible to everyone, not just trading firms. Use my access code [abcd-1234-efgh-5678] to skip the waitlist and start building your first automated strategy today.\n\nHey -- I've been quiet about what we're building at [Plug](/) for the past year. Today, that changes.\n\nIf you're like me, you've probably missed countless opportunities because you were sleeping, in a meeting, or simply away from your computer. The reality is that crypto never sleeps, but we're not machines – we need rest, we have jobs, and we have lives outside of watching charts and Discord announcements.\n\nThat's why we built Plug. It's your always-on presence in the digital world, executing your strategies exactly as you'd do them yourself. No more missed airdrops, no more liquidations because you couldn't react in time, and no more manual compounding of your DeFi positions.\n\n## Where We Are Today\n\nToday marks an important milestone as we open Plug's beta to everyone. From day one, you'll be able to automate your onchain presence across major networks like Ethereum, Optimism, Base, and Arbitrum. Connect with the protocols you already use through our simple drag-and-drop interface, set up custom triggers based on price movements or time schedules, and monitor everything in real-time.\n\nWhat excites me most is how we've made complex blockchain interactions feel natural. You don't need to understand the underlying mechanics – just tell Plug what you want to achieve, and it figures out the best way to make it happen.\n\n## The Road Ahead\n\nThis is just the beginning. We're already working with major protocols to expand our integration suite. Soon, you'll be able to create even more sophisticated automations, share your successful strategies with others (and yes, earn from them), and manage everything from your phone.\n\nThe vision is clear: we want to make onchain automation accessible to everyone, not just trading firms with million-dollar budgets. Whether you're looking to DCA into ETH, auto-compound your staking rewards, or bid on NFTs with specific traits, Plug will handle it all.\n\n## Join Us\n\nWe're building Plug in public because we believe the best products are shaped by their users. Start small – maybe set up a simple DCA strategy or automate your reward compounding. Tell us what works, what doesn't, and what you wish you could do.\n\nI've created a special access code for early readers: [abcd-1234-efgh-5678]. Use it at [app.onplug.io](https://app.onplug.io) to skip the waitlist and start building your first automation.\n\nFollow our journey:\n\n- Twitter: [[@onplug_io]](https://twitter.com/onplug_io)\n- Documentation: [docs.onplug.io](https://docs.onplug.io)\n\nSee you onchain,\nChance\n",attributes:{created:"2025-01-10T06:00:00.000Z",className:"",tags:["perspective"],author:"nftchance"}}};

        // * Get all the Posts for a given page.
        export const getPosts = (
            page = 1,
            pageSize = PAGE_SIZE,
            filter?:
                | Partial<Record<'date' | 'tag' | 'search', Partial<string>>>
                | undefined
        ): {
            posts: Post[]
            count: number
            hasNext: boolean
            random: Post
        } => {
            // ! Filter the posts before paginating so that we can get a final count.
            const filteredPosts = Object.values(
                Object.values(posts ?? {}).reduce<Record<string, Post>>(
                    (acc, article) => {
                        if(article.attributes.unlisted) return acc

                        if (!filter) {
                            acc[article.slug] = article

                            return acc
                        }

                        const { tag, date, search } = filter

                        // * Check the lowercase tags against the lowercase tag parameter.
                        const matchesTag =
                            !tag ||
                            (article.attributes.tags?.some(
                                articleTag =>
                                    articleTag.toLowerCase() === tag.toLowerCase()
                            ) ??
                                false)

                        // ! Require the exact date match.
                        const matchesDate =
                            !date ||
                            ((article.attributes.created?.includes(date) ||
                                article.attributes.updated?.includes(date)) ??
                                false)

                        const matchesSearch =
                            !search ||
                            article.title
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            article.description
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            article.content
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            (article.attributes.tags?.some(articleTag =>
                                articleTag.toLowerCase().includes(search.toLowerCase())
                            ) ??
                                false)

                        if (matchesTag && matchesDate && matchesSearch) {
                            acc[article.slug] = article
                        }

                        return acc
                    },
                    {}
                )
            )

            const count = filteredPosts.length
            const keys = Object.keys(posts)

            return {
                posts: filteredPosts.slice(
                    (page - 1) * pageSize,
                    page * pageSize
                ),
                count,
                hasNext: count > page * pageSize,
                random: posts[keys[Math.floor(Math.random() * keys.length)]]
            }
        }

        // * Get a specific Post by the value of the slug parameter in each Post.
        // ! The parameter can be the dictionary key or the slug which is a slug on the Post.
        export type PostLookupKey = keyof typeof posts
        export type PostLookup = (typeof posts)[PostLookupKey]['slug']

        export const getPost = (lookup: PostLookup) => {
            const article = posts[lookup.replaceAll('-', '') as PostLookupKey]

            if (!article) throw new Error('Post not found')

            return article
        }

        // * Get the favicon for a given URL.
        export const getFavicon = (url?: string) => {
            if (!url) return null

            const faviconUrl =
                faviconUrls[
                    url
                        .replace('https://', '')
                        .replace('http://', '') as keyof typeof faviconUrls
                ]

            if (!faviconUrl) return null

            return faviconUrl
        }
        